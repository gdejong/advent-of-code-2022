package year2022

import "testing"

var testInputDay02 = []string{
	"A Y",
	"B X",
	"C Z",
}

func TestDay02_PartA(t *testing.T) {
	day := Day02{}

	got := day.PartA(testInputDay02)

	want := 15
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestDay02_PartB(t *testing.T) {
	day := Day02{}

	got := day.PartB(testInputDay02)

	want := 12
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
