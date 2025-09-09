package main

import "fmt"

func function1() {
	fmt.Println("I must be the first function to run")
}

func function2() {
	fmt.Println("I must be the second function to run")
}

func function3() {
	fmt.Println("I must be the third function to run")
}

func function4() {
	fmt.Println("I must be the fourth function to run")
}

func function5() {
	fmt.Println("I must be the fifth function to run")
}

func main() {

	function1()
	defer function2()
	function3()
	defer function4()
	function5()
}

// As you can see, the 'defer' statement respect the LIFO order (Last In First Out)
// That's why 'function4()' run first than 'function2()'.
