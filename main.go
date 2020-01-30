package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Initialize data structures
	var cpu = CPU{}
	var memory = make(Memory)
	var stack Stack

	// Load the program from memory and get location of first instruction
	firstInstruction, err := memory.LoadProgram("./input.txt")
	if err != nil {
		panic(err)
	}
	cpu.PC = firstInstruction

	// Run until the program ends
	for err == nil {
		err = cpu.ExecuteNext(&memory, &stack)
	}
}

// Convenience functions
func hexToDec(hex string) uint16 {
	dec, err := strconv.ParseUint(hex, 16, 16)
	if err != nil {
		panic(err)
	}
	return uint16(dec)
}

func decToHex(dec uint16) string {
	return fmt.Sprintf("%x", dec)
}
