package machine

import (
	"github.com/Velnbur/AsmVM/register"
	"testing"
)

func TestMachineParseOperator(t *testing.T) {
	var registers [CountRegs]register.Register

	for i := 0; i < CountRegs; i++ {
		registers[i] = register.New()
	}

	machine := state{
		Registers:    registers,
		Command:      "",
		TactsCounter: 0,
		CommCounter:  0,
		Sign:         false,
		operator:     NONE,
		operand1:     nil,
		operand2:     nil,
	}
	machine.parseOperator("mov")

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
		Registers:    registers,
		Command:      "",
		TactsCounter: 0,
		CommCounter:  0,
		Sign:         false,
		operator:     NONE,
		operand1:     nil,
		operand2:     nil,
	}
	machine.parseOperand("R2")

	if machine.operand1 != machine.Registers[1] {
		t.Fatalf("Defining R2 failed")
	}
	machine.operand1 = nil
	machine.parseOperand("12")

	if machine.operand1.Value() != 12 {
		t.Fatalf("Defining 12 failed")
	}

	err := machine.parseOperand("R10")
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
		Registers:    registers,
		Command:      "",
		TactsCounter: 0,
		CommCounter:  0,
		Sign:         false,
		operator:     NONE,
		operand1:     nil,
		operand2:     nil,
	}

	machine.Command = "mov R1, R2"
	err := machine.parseLine()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if machine.operator != MOV {
		t.Fatalf("Operator must be 'MOV' ")
	}
	if machine.operand1 != machine.Registers[0] {
		t.Fatalf("First operand must be R1")
	}
	if machine.operand2 != machine.Registers[1] {
		t.Fatalf("Second operand must be R2")
	}
}

func TestHandleCommand(t *testing.T) {
	var registers [CountRegs]register.Register

	for i := 0; i < CountRegs; i++ {
		registers[i] = register.New()
	}

	machine := state{
		Registers:    registers,
		Command:      "",
		TactsCounter: 0,
		CommCounter:  0,
		Sign:         false,
		operator:     NONE,
		operand1:     nil,
		operand2:     nil,
	}

	machine.Command = "mov R1, R2"

	err := machine.HandleCommand("mov R1, 12")
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if machine.Registers[0].Value() != 12 {
		t.Fatalf("R1 must be 12")
	}

	machine.HandleCommand("mov R2, 11")
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if machine.Registers[1].Value() != 11 {
		t.Fatalf("R2 must be 11")
	}

	machine.HandleCommand("sub R1, R2")
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if machine.Registers[0].Value() != 1 && machine.Sign != true {
		t.Fatalf("R1!=1 and Sign!=true")
	}
}
