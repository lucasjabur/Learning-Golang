package main

import "testing"

func TestFactorial(t *testing.T) {
	factorial_result := factorial(6)
	expected_result := 720
	if factorial_result != expected_result {
		t.Errorf("Factorial result was incorrect, got: %d, want: %d\n", factorial_result, expected_result)
	}
}
