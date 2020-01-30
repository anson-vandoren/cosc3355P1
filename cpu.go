package main

import (
	"errors"
	"fmt"
)

// A CPU represents the hypothetical machine from page 6
type CPU struct {
	AC  uint16
	PC  uint16
	IR  uint16
	REG uint16
}

// ExecuteNext loads the next instruction, executes it, and increments the program counter
func (cpu *CPU) ExecuteNext(memory *Memory, stack *Stack) (err error) {
	// Load next instruction into the Instruction Register
	cpu.IR = (*memory)[cpu.PC]
	// Increment the program counter
	cpu.PC++

	// Parse the instruction
	opcode := cpu.IR >> 12           // four bits for opcode
	addr := hexToDec("FFF") & cpu.IR // remaining twelve bits for address

	// Execute the opcode
	switch opcode {
	case 0:
		// No-op
	case 1:
		// Load AC from memory
		cpu.AC = (*memory)[addr]
	case 2:
		// Store AC to memory
		(*memory)[addr] = cpu.AC
	case 3:
		// Load AC from REG
		cpu.AC = cpu.REG
	case 4:
		// Store AC to REG
		cpu.REG = cpu.AC
	case 5:
		// Add to AC from memory
		cpu.AC += (*memory)[addr]
	case 6:
		// Load REG with operand
		cpu.REG = addr
	case 7:
		// Add REG to AC
		cpu.AC += cpu.REG
	case 8:
		// Multiply REG into AC
		cpu.AC *= cpu.REG
	case 9:
		// Subtract REG from AC
		cpu.AC -= cpu.REG
	case 10:
		// Divide AC by REG (integer division)
		cpu.AC /= cpu.REG
	case 11:
		// Save registers to stack
		stack.Push(cpu.AC)
		stack.Push(cpu.PC)
		stack.Push(cpu.IR)
		stack.Push(cpu.REG)
		// Jump to subroutine starting at the address
		cpu.PC = addr
	case 12:
		// Return from subroutine
		// Restore registers from stack
		cpu.REG = stack.Pop()
		cpu.IR = stack.Pop()
		cpu.PC = stack.Pop()
		cpu.AC = stack.Pop()
		fmt.Printf("After popping from stack:\nREG=%X\nIR=%X\nPC=%X\nAC=%X\n", cpu.REG, cpu.IR, cpu.PC, cpu.AC)
	case 15:
		// Halt (end of program)
		err = errors.New("Program ended")
		fmt.Printf("At end, AC=%X, REG=%X, PC=%X, IR=%X, 940=%X, 941=%X, 942=%X\n\n", cpu.AC, cpu.REG, cpu.PC, cpu.IR, memory.Load("940"), memory.Load("941"), memory.Load("942"))
	}
	fmt.Printf("Next instruction location is %X and err is %v\n", cpu.PC, err)
	return err
}
