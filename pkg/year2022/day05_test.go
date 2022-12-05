package year2022

import "testing"

var testInputDay05 = []string{
	"    [D]    ",
	"[N] [C]    ",
	"[Z] [M] [P]",
	" 1   2   3 ",
	"",
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

func TestDay05_PartA(t *testing.T) {
	day := Day05{}

	got := day.PartA(testInputDay05)

	want := "CMZ"
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestDay05_PartB(t *testing.T) {
	day := Day05{}

	got := day.PartB(testInputDay05)

	want := "MCD"
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
