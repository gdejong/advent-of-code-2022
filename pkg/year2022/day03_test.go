package year2022

import "testing"

var testInputDay03 = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestDay03_PartA(t *testing.T) {
	day := Day03{}

	got := day.PartA(testInputDay03)

	want := 157
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestDay03_PartB(t *testing.T) {
	day := Day03{}

	got := day.PartB(testInputDay03)

	want := 70
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
