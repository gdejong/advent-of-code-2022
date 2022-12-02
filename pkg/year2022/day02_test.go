package year2022

import "testing"

func TestDay02_PartA(t *testing.T) {
	day := Day02{}

	testInput := []string{
		"A Y",
		"B X",
		"C Z",
	}

	got := day.PartA(testInput)

	want := 15
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestDay02_PartB(t *testing.T) {
	day := Day02{}

	testInput := []string{
		"A Y",
		"B X",
		"C Z",
	}

	got := day.PartB(testInput)

	want := 12
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
