package year2022

import "testing"

var testInputDay04 = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func TestDay04_PartA(t *testing.T) {
	day := Day04{}

	got := day.PartA(testInputDay04)

	want := 2
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestDay04_PartB(t *testing.T) {
	day := Day04{}

	got := day.PartB(testInputDay04)

	want := 4
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
