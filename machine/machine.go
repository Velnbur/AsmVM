package machine

import (
	"errors"
	"fmt"
	reg "github.com/Velnbur/AsmVM/register"
	"github.com/Velnbur/AsmVM/token"
	"strconv"
	"strings"
)

const CountRegs = 8

type OPERAND uint

const (
	MOV OPERAND = iota
	SUB
	NONE
)

type Machine interface {
	// PrintStatus prints current state
	PrintStatus()
	HandleCommand(line string) error
}

type state struct {
	Registers    [CountRegs]reg.Register
	CommCounter  int
	TactsCounter int
	Sign         bool // result less then zero if true
	Command      string
	operator     OPERAND
	operand1     token.Token
	operand2     token.Token
}

func New() Machine {
	var registers [CountRegs]reg.Register

	for i := 0; i < CountRegs; i++ {
		registers[i] = reg.New()
	}

	return &state{
		Registers:    registers,
		Command:      "",
		TactsCounter: 0,
		CommCounter:  0,
		Sign:         false,
		operator:     NONE,
		operand1:     nil,
		operand2:     nil,
	}
}

func (m state) PrintStatus() {
	fmt.Printf("IR: %s\n", m.Command)
	for i, reg := range m.Registers {
		s := reg.ToSlice()
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		fmt.Printf("R%d: %v\n", i+1, s)
	}
	fmt.Printf("PS: %t\n", m.Sign)
	fmt.Printf("PC: %d\n", m.CommCounter)
	fmt.Printf("TC: %d\n\n", m.TactsCounter)
}

func (m *state) parseOperator(oper string) error {
	switch oper {
	case "mov":
		m.operator = MOV
		return nil
	case "sub":
		m.operator = SUB
		return nil
	}
	return errors.New("No such operator")
}

// parse string and fill last empty operand
// in machine with Token
// Token can be Value or Register type
func (m *state) parseOperand(oper string) error {
	if oper[0] == 'R' {
		if len(oper) != 2 {
			return errors.New("No such register")
		}
		if i := oper[1] - '0'; i < CountRegs-1 && i > 0 {
			if m.operand1 == nil {
				m.operand1 = m.Registers[i-1]
			} else if m.operand2 == nil {
				m.operand2 = m.Registers[i-1]
			}
			return nil
		} else {
			return errors.New("No such register")
		}
	}

	num, err := strconv.Atoi(oper)
	if err != nil {
		return err
	}


	if m.operand1 == nil {
		m.operand1 = token.New(int16(num))
	} else if m.operand2 == nil {
		m.operand2 = token.New(int16(num))
	}

	return nil
}

func (m *state) parseLine() error {
	// temp slice for operands and operator
	s := strings.Split(strings.Replace(m.Command, ",", "", -1), " ")
	if len(s) != 3 {
		return errors.New("Syntax error")
	}

	err := m.parseOperator(s[0])
	err = m.parseOperand(s[1])
	err = m.parseOperand(s[2])
	if err != nil {
		return err
	}

	return nil
}

func (m *state) Clean() {
	m.operand1 = nil
	m.operand2 = nil
	m.TactsCounter = 0
	m.Sign = false
}

func (m *state) HandleCommand(line string) error {
	m.Clean()
	m.Command = line
	m.CommCounter++

	err := m.parseLine()
	if err != nil {
		return err
	}
	if m.operator == MOV {
		m.mov()
	} else if m.operator == SUB {
		m.sub()
	}
	return nil
}

func (m *state) mov() {
	switch m.operand1.(type) {
	case reg.Register:
		m.operand1.Change(m.operand2.Value())
	}
}

func (m *state) sub() {
	result := m.operand1.Value() - m.operand2.Value()
	if result < 0 {
		m.Sign = true
	}

	m.Registers[0].Change(result)
}
