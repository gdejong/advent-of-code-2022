package year2022

import "testing"

func TestDay03_PartA(t *testing.T) {
	day := Day03{}

	testInput := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	got := day.PartA(testInput)

	want := 157
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestDay03_PartB(t *testing.T) {
	day := Day03{}

	testInput := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	got := day.PartB(testInput)

	want := 70
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
