package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()
	if s.Len() != 0 {
		t.Errorf("size should be 0")
	}
	s.Push("x")
	if s.Len() != 1 {
		t.Errorf("size should be 1")
	}
	s.Push("y")
	if s.Len() != 2 {
		t.Errorf("size should be 2")
	}

	if s.Pop() != "y" {
		t.Errorf("incorrect data popped")
	}
	if s.Len() != 1 {
		t.Errorf("size should be 0")
	}

	if s.Pop() != "x" {
		t.Errorf("incorrect data popped")
	}
	if s.Len() != 0 {
		t.Errorf("size should be 0")
	}
}
