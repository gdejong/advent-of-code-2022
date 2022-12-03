package year2022

import (
	"gdejong/advent-of-code-2022/pkg/input"
	"strings"
	"unicode"
)

type Day03 struct{}

func findOverlap(a, b string) rune {
	for _, c := range a {
		if strings.ContainsRune(b, c) {
			return c
		}
	}

	panic("no overlap found")
}

func convertToPriority(char rune) int {
	// Lowercase item types a through z have priorities 1 through 26
	// ASCII value for 'a' is 97
	if unicode.IsLower(char) {
		return int(char - 96)
	}

	// Uppercase item types A through Z have priorities 27 through 52
	// ASCII value for 'A' is 65
	return int(char - 38)
}

func (p Day03) PartA(lines []string) any {
	fn := func(line string) int {
		compartmentA := line[0 : len(line)/2] // first half of string
		compartmentB := line[len(line)/2:]    // second half of string

		overlappingRune := findOverlap(compartmentA, compartmentB)

		return convertToPriority(overlappingRune)
	}

	return input.ArrayMapAndSumOverNonEmptyLines(lines, fn)
}

func threeWayOverlap(group []string) rune {
	for _, c := range group[0] {
		if strings.ContainsRune(group[1], c) && strings.ContainsRune(group[2], c) {
			return c
		}
	}

	for _, c := range group[1] {
		if strings.ContainsRune(group[0], c) && strings.ContainsRune(group[2], c) {
			return c
		}
	}

	for _, c := range group[2] {
		if strings.ContainsRune(group[0], c) && strings.ContainsRune(group[1], c) {
			return c
		}
	}

	panic("no overlap found")
}

func (p Day03) PartB(lines []string) any {
	groups := input.Chunk(lines, 3)

	total := 0
	for _, group := range groups {
		overlappingRune := threeWayOverlap(group)
		total += convertToPriority(overlappingRune)
	}

	return total
}
