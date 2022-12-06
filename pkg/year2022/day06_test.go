package year2022

import "testing"

var testInputDay06 = []string{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
}

func TestDay06_PartA(t *testing.T) {
	day := Day06{}

	got := day.PartA(testInputDay06)

	want := 7
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestDay06_PartB(t *testing.T) {
	day := Day06{}

	got := day.PartB(testInputDay06)

	want := 19
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
