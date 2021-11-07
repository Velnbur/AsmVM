package machine

import (
	"errors"
	"fmt"
	reg "github.com/Velnbur/AsmVM/src/register"
	"github.com/Velnbur/AsmVM/src/token"
	"strconv"
	"strings"
)

// CountRegs - amount of registers in machine
const CountRegs = 8

// Operator - just to mark operators
// in lower enum
type Operator uint

const (
	MOV Operator = iota
	SUB
	NONE
)

// Machine - represents VM object
// that reacts to all commands from source file
// and changes its inner state
type Machine interface {
	// SetCommand - set new command that will be
	// handled. Always must be called before HandleCommand
	SetCommand(line string) error

	// PrintStatus prints current
	// state parameters to STDOUT
	PrintStatus()

	// HandleCommand - react to command and change
	// inner state. Return error if syntax
	// of command was incorrect
	HandleCommand() error
}

// state struct - represents current state of VM
type state struct {
	registers   [CountRegs]reg.Register
	commCounter int
	tactsCounter int
	sign         bool // result was less than zero if true
	command     string
	operator    Operator
	operand1    token.Token
	operand2 token.Token
}

// New - Machine constructor
func New() Machine {
	var registers [CountRegs]reg.Register

	for i := 0; i < CountRegs; i++ {
		registers[i] = reg.New()
	}

	return &state{
		registers:    registers,
		command:      "",
		tactsCounter: 0,
		commCounter:  0,
		sign:         false,
		operator:     NONE,
		operand1:     nil,
		operand2:     nil,
	}
}

func (m *state) SetCommand(line string) error {
	m.clean()
	m.tactsCounter = 1
	m.command = line
	err := m.parseLine()
	if err != nil {
		return err
	}
	return nil
}

func (m state) PrintStatus() {
	fmt.Printf("IR: %s\n", m.command)
	for l, register := range m.registers {
		s := register.ToSlice()
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		fmt.Printf("R%d: ", l+1)
		for i := 0; i < len(s); i++ {
			fmt.Print(s[i])
		}
		fmt.Println()
	}
	fmt.Print("PS: ")
	if m.sign {
		fmt.Print("1")
	} else {
		fmt.Print("0")
	}
	fmt.Printf("\nPC: %d\n", m.commCounter)
	fmt.Printf("TC: %d\n", m.tactsCounter)
}

// parseOperator - parse operator from string
// change operator field in state struct to one
// of Operator enum
func (m *state) parseOperator(oper string) error {
	switch oper {
	case "mov":
		m.operator = MOV
		return nil
	case "sub":
		m.operator = SUB
		return nil
	}
	return errors.New("no such operator")
}

// parseOperand - parse string and fill last empty operand
// in state struct with Token.
// Token can be SimpleValue or Register type
func (m *state) parseOperand(operand string) error {
	if operand[0] == 'R' {
		if len(operand) != 2 {
			return errors.New("no such register")
		}
		if i := operand[1] - '0'; i < CountRegs-1 && i > 0 {
			if m.operand1 == nil {
				m.operand1 = m.registers[i-1]
			} else if m.operand2 == nil {
				m.operand2 = m.registers[i-1]
			} else {
				return errors.New("all operands are filled")
			}
			return nil
		} else {
			return errors.New("no such register")
		}
	}

	num, err := strconv.Atoi(operand)
	if err != nil {
		return err
	}


	if m.operand1 == nil {
		m.operand1 = token.NewSimpleValue(int16(num))
	} else if m.operand2 == nil {
		m.operand2 = token.NewSimpleValue(int16(num))
	} else {
		return errors.New("all operands are filled")
	}

	return nil
}

// parseLine - parses command field if state struct
// into Operator and two operands (Tokens).
// Call parseOperator, parseOperand
func (m *state) parseLine() error {
	// temp slice for operands and operator
	s := strings.Split(strings.Replace(m.command, ",", "", -1), " ")
	if len(s) != 3 {
		return errors.New("syntax error")
	}

	err := m.parseOperator(s[0])
	err = m.parseOperand(s[1])
	err = m.parseOperand(s[2])
	if err != nil {
		return err
	}

	return nil
}

// clean - clean used fields in state
// before starting HandleCommand
func (m *state) clean() {
	m.operator = NONE
	m.operand1 = nil
	m.operand2 = nil
	m.sign = false
	m.command = ""
}

func (m *state) HandleCommand() error {
	m.tactsCounter++
	m.commCounter++

	switch m.operator {
	case MOV:
		m.mov()
	case SUB:
		m.sub()
	default:
		return errors.New("wrong operator")
	}

	return nil
}

// mov - 'mov' operator method
func (m *state) mov() {
	switch m.operand1.(type) {
	case reg.Register:
		m.operand1.Change(m.operand2.Value())
	}
}

// sub - 'sub' operator method
func (m *state) sub() {
	result := m.operand1.Value() - m.operand2.Value()
	if result < 0 {
		m.sign = true
	}

	for _, register := range m.registers {
		if register == m.operand1 {
			register.Change(result)
			break
		}
	}
}
