package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

// Inside the directory of this file, run: 'go test'
// Output:
/*

	PASS
	ok      mymodule/section20/hands_on_exercise_63 0.002s

*/
