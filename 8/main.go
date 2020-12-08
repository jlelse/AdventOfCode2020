package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type instruction struct {
	op       string
	arg      int
	executed bool
}

func main() {
	instructionRe := regexp.MustCompile(`^(?P<op>\w+) (?P<arg>[\+\-]\d+)$`)
	parsedInstructions := []instruction{}

	file, _ := os.Open("input.txt")
	defer file.Close()
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		text := scanner.Text()
		arg, _ := strconv.Atoi(instructionRe.ReplaceAllString(text, "$arg"))
		parsedInstructions = append(parsedInstructions, instruction{instructionRe.ReplaceAllString(text, "$op"), arg, false})
	}

	makeInstructions := func() []instruction {
		instructions := make([]instruction, len(parsedInstructions))
		copy(instructions, parsedInstructions)
		return instructions
	}

	var accumulator, pos int

	for newInstructions := makeInstructions(); true; {
		instr := &newInstructions[pos]
		if instr.executed {
			break
		}
		switch instr.op {
		case "acc":
			accumulator += instr.arg
			pos++
		case "nop":
			pos++
		case "jmp":
			pos += instr.arg
		}
		instr.executed = true
	}

	fmt.Println("Part 1:", accumulator)

modifyOpLoop:
	for lastModifiedOp := 0; true; lastModifiedOp++ {
		newInstructions := makeInstructions()
		if instr := &newInstructions[lastModifiedOp]; instr.op == "jmp" {
			instr.op = "nop"
		} else if instr.op == "nop" {
			instr.op = "jmp"
		} else {
			continue
		}
		accumulator = 0
		pos = 0
		for {
			if pos >= len(newInstructions) {
				break modifyOpLoop
			}
			instr := &newInstructions[pos]
			if instr.executed {
				break
			}
			switch instr.op {
			case "acc":
				accumulator += instr.arg
				pos++
			case "nop":
				pos++
			case "jmp":
				pos += instr.arg
			}
			instr.executed = true
		}
	}

	fmt.Println("Part 2:", accumulator)
}
