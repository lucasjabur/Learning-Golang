package main

func Add(a int, b int) int {
	return a + b
}

// A little bit about conventions:
/*
	1. Test files: Test files live in the same package as the code they test. They are named with
	the following convention: `filename_test.go` where filename is the name of the file that
	contains the code you want to test. In this case, 'main_test.go'

	2. Test functions: Test functions must start with the word `Test` followed by a word that starts
	with a capital letter. The signature of a test function is `func TestXxx(*testing.T)`, where Xxx
	does not start with a lowercase letter. In this case, 'func TestAdd(t *testing.T)'
*/
