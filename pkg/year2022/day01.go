package year2022

import (
	"sort"
	"strconv"
)

type Day01 struct{}

func (p Day01) PartA(lines []string) any {
	caloriesPerElf := countCalories(lines)

	return caloriesPerElf[len(caloriesPerElf)-1]
}

func (p Day01) PartB(lines []string) any {
	caloriesPerElf := countCalories(lines)

	return caloriesPerElf[len(caloriesPerElf)-1] + caloriesPerElf[len(caloriesPerElf)-2] + caloriesPerElf[len(caloriesPerElf)-3]
}

func countCalories(input []string) []int {
	elves := []int{0}
	currentElf := 0
	for _, line := range input {
		if line == "" {
			currentElf++
			elves = append(elves, 0)
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		elves[currentElf] += calories
	}

	sort.Ints(elves)

	return elves
}
