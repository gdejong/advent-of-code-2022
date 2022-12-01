package year2022

import "testing"

func TestDay01_PartA(t *testing.T) {
	testInput := []string{
		"1",
		"",
		"1",
		"1",
		"",
		"1",
		"1",
		"1",
	}
	day := Day01{}
	got := day.PartA(testInput)
	want := 3
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestDay01_PartB(t *testing.T) {
	testInput := []string{
		"1",
		"",
		"1",
		"1",
		"",
		"1",
		"1",
		"1",
	}
	day := Day01{}
	got := day.PartB(testInput)
	want := 6
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
