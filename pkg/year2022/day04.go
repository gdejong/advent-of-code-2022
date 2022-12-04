package year2022

import (
	"gdejong/advent-of-code-2022/pkg/input"
	"strings"
)

type Day04 struct{}

func (p Day04) PartA(lines []string) any {
	return input.ArrayMapAndSumOverNonEmptyLines(lines, func(line string) int {
		split := strings.Split(line, ",")
		first, second := split[0], split[1]

		firstSections := input.StringSliceToIntSlice(strings.Split(first, "-"))
		secondSections := input.StringSliceToIntSlice(strings.Split(second, "-"))

		if (firstSections[0] >= secondSections[0] && firstSections[1] <= secondSections[1]) || (secondSections[0] >= firstSections[0] && secondSections[1] <= firstSections[1]) {
			return 1
		}

		return 0
	})
}

func (p Day04) PartB(lines []string) any {
	return input.ArrayMapAndSumOverNonEmptyLines(lines, func(line string) int {
		split := strings.Split(line, ",")
		first, second := split[0], split[1]

		firstSections := input.StringSliceToIntSlice(strings.Split(first, "-"))
		secondSections := input.StringSliceToIntSlice(strings.Split(second, "-"))

		if firstSections[1] >= secondSections[0] && firstSections[0] <= secondSections[1] {
			return 1
		}

		return 0
	})
}
