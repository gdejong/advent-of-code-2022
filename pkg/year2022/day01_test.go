package year2022

import "testing"

var testInputDay01 = []string{
	"1000",
	"2000",
	"3000",
	"",
	"4000",
	"",
	"5000",
	"6000",
	"",
	"7000",
	"8000",
	"9000",
	"",
	"10000",
}

func TestDay01_PartA(t *testing.T) {
	day := Day01{}

	got := day.PartA(testInputDay01)

	want := 24000
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestDay01_PartB(t *testing.T) {
	day := Day01{}

	got := day.PartB(testInputDay01)

	want := 45000
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
