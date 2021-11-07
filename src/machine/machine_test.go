package machine

import (
	"github.com/Velnbur/AsmVM/src/register"
	"testing"
)

func TestMachineParseOperator(t *testing.T) {
	var registers [CountRegs]register.Register

	for i := 0; i < CountRegs; i++ {
		registers[i] = register.New()
	}

	machine := state{
		registers:    registers,
		command:      "",
		tactsCounter: 0,
		commCounter:  0,
		sign:         false,
		operator:     NONE,
		operand1:     nil,
		operand2:     nil,
	}
	err := machine.parseOperator("mov")
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if machine.operator != MOV {
		t.Fatalf("Parse 'mov' failed, must be MOV")
	}
}

func TestMachineParseOperand(t *testing.T) {
	var registers [CountRegs]register.Register

	for i := 0; i < CountRegs; i++ {
		registers[i] = register.New()
	}

	machine := state{
		registers:    registers,
		command:      "",
		tactsCounter: 0,
		commCounter:  0,
		sign:         false,
		operator:     NONE,
		operand1:     nil,
		operand2:     nil,
	}
	err := machine.parseOperand("R2")
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if machine.operand1 != machine.registers[1] {
		t.Fatalf("Defining R2 failed")
	}
	machine.operand1 = nil
	err = machine.parseOperand("12")
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if machine.operand1.Value() != 12 {
		t.Fatalf("Defining 12 failed")
	}

	err = machine.parseOperand("R10")
	if err == nil {
		t.Fatalf("There is no such register as R10")
	}
}

func TestMachineParseLine(t *testing.T) {
	var registers [CountRegs]register.Register

	for i := 0; i < CountRegs; i++ {
		registers[i] = register.New()
	}

	machine := state{
		registers:    registers,
		command:      "",
		tactsCounter: 0,
		commCounter:  0,
		sign:         false,
		operator:     NONE,
		operand1:     nil,
		operand2:     nil,
	}

	machine.command = "mov R1, R2"
	err := machine.parseLine()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if machine.operator != MOV {
		t.Fatalf("Operator must be 'MOV' ")
	}
	if machine.operand1 != machine.registers[0] {
		t.Fatalf("First operand must be R1")
	}
	if machine.operand2 != machine.registers[1] {
		t.Fatalf("Second operand must be R2")
	}
}

func TestHandleCommand(t *testing.T) {
	var registers [CountRegs]register.Register

	for i := 0; i < CountRegs; i++ {
		registers[i] = register.New()
	}

	machine := state{
		registers:    registers,
		command:      "",
		tactsCounter: 0,
		commCounter:  0,
		sign:         false,
		operator:     NONE,
		operand1:     nil,
		operand2:     nil,
	}

	machine.command = "mov R1, R2"

	machine.SetCommand("mov R1, 12")
	err := machine.HandleCommand()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if machine.registers[0].Value() != 12 {
		t.Fatalf("R1 must be 12")
	}

	machine.SetCommand("mov R2, 11")
	err = machine.HandleCommand()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if machine.registers[1].Value() != 11 {
		t.Fatalf("R2 must be 11")
	}

	machine.SetCommand("sub R1, R2")
	err = machine.HandleCommand()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if machine.registers[0].Value() != 1 && machine.sign != true {
		t.Fatalf("R1!=1 and sign!=true")
	}
}
