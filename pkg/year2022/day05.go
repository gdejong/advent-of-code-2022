package year2022

import (
	"gdejong/advent-of-code-2022/pkg/input"
	"gdejong/advent-of-code-2022/pkg/stack"
	"strings"
)

type Day05 struct{}

type instruction struct {
	howMany   int
	fromIndex int
	toIndex   int
}

func instructions(lines []string) []instruction {
	var instructions []instruction

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "move") {
			parsed := strings.Fields(line)
			instructions = append(instructions, instruction{
				howMany:   input.MustInt(parsed[1]),
				fromIndex: input.MustInt(parsed[3]) - 1,
				toIndex:   input.MustInt(parsed[5]) - 1,
			})
		}
	}

	return instructions
}

func parseStacks(lines []string) []*stack.Stack {
	// Create stacks
	var stacks []*stack.Stack
	c := len(input.ChunkString(lines[0], 4)) // look at first line to determine amount of stacks needed
	for i := 0; i < c; i++ {
		stacks = append(stacks, stack.New())
	}

	for _, line := range lines {
		// Stop parsing stacks when the line with numbers comes up
		if strings.Fields(line)[0] == "1" {
			return stacks
		}

		c := input.ChunkString(line, 4)
		for i, chucked := range c {
			if strings.TrimSpace(chucked) != "" {
				letter := strings.TrimSpace(chucked)[1:2] // From "[D]" to "D"
				stacks[i].Unshift(letter)
			}
		}
	}

	panic("oops")
}

func (p Day05) PartA(lines []string) any {
	stacks := parseStacks(lines)
	instructionsToPerform := instructions(lines)

	for _, instructionToPerform := range instructionsToPerform {
		for i := 0; i < instructionToPerform.howMany; i++ {
			stacks[instructionToPerform.toIndex].Push(stacks[instructionToPerform.fromIndex].Pop())
		}
	}

	answer := ""
	for _, s := range stacks {
		answer += s.Peek()
	}
	return answer
}

func (p Day05) PartB(lines []string) any {
	stacks := parseStacks(lines)
	instructionsToPerform := instructions(lines)

	for _, instructionToPerform := range instructionsToPerform {
		var toPutBack []string
		for i := 0; i < instructionToPerform.howMany; i++ {
			toPutBack = append(toPutBack, stacks[instructionToPerform.fromIndex].Pop())
		}

		for i := len(toPutBack) - 1; i >= 0; i-- {
			stacks[instructionToPerform.toIndex].Push(toPutBack[i])
		}
	}

	answer := ""
	for _, s := range stacks {
		answer += s.Peek()
	}
	return answer
}
