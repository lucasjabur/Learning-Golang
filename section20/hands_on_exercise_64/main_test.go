package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(10, 10)
	if total != 20 {
		t.Errorf("Sum was incorrect, got: %d, want: %d\n", total, 20)
	}
}

func TestSubtract(t *testing.T) {
	total := subtract(10, 5)
	if total != 5 {
		t.Errorf("Sum was incorrect, got: %d, want: %d\n", total, 5)
	}
}

func TestDoMath(t *testing.T) {
	total := doMath(42, 16, add)
	if total != 58 {
		t.Errorf("Math operation was incorrect, got: %d, want: %d\n", total, 58)
	}
}
