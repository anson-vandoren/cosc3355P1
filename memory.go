package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Memory represents a block of 16-bit memory locations
type Memory map[uint16]uint16

// Load returns the value stored in memory at the given address
func (m Memory) Load(addr string) uint16 {
	intAddr := hexToDec(addr)
	return m[uint16(intAddr)]
}

// Store places the given value at the memory location given by addr
// where addr is a hex string
func (m Memory) Store(addr string, val uint16) (err error) {
	intAddr, err := strconv.ParseInt(addr, 16, 16)
	if err != nil {
		return err
	}
	m[uint16(intAddr)] = val
	return nil
}

// LoadProgram reads in a program text file and stores it in memory
func (m Memory) LoadProgram(filename string) (start uint16, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Match decimal line numbers, and all comments
	stripLines := regexp.MustCompile(`(^\d+\.\s+|\s*;.+$)`)

	for scanner.Scan() {
		// Read next line and trim whitespace from both ends
		nextLine := strings.TrimSpace(scanner.Text())
		// Ignore comments
		if nextLine[0] == '=' {
			continue
		}
		// Remove decimal line numbers and all comments
		bareLine := stripLines.ReplaceAllString(nextLine, "")

		// Separate memory location and instruction for that address
		tokens := strings.Split(bareLine, " ")
		if len(tokens) != 2 {
			err = fmt.Errorf("Expected 2 tokens per line but got %d", len(tokens))
			return
		}

		// Store the instruction at the memory address
		fmt.Printf("Address: %s, Instruction: %s\n", tokens[0], tokens[1])
		err = m.Store(tokens[0], hexToDec(tokens[1]))

		// Set the first address encountered as the starting point
		if start == 0 {
			start = hexToDec(tokens[0])
		}
	}
	return
}
