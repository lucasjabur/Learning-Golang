package main

import "fmt"

func normalFunction() {
	fmt.Println("I'm not an anonymous function!")
}

func main() {

	normalFunction()

	func() { // func(){}() --> func (parameters) { code } (arguments), can have 'returns' as well!
		fmt.Println("I am an anonymous function!")
	}()

	func(name string) {
		fmt.Println("This is an anonymous function showing my name:", name)
	}("Lucas")

	// You can assign functions to variables.
	// Functions, values and types are 'first class citizens' in Golang, this means that they have the same rights, let's say.
	// So, you could do:

	function1 := func() {
		fmt.Println("Function1 is printing...")
	}
	function1()

	function2 := func(language string) {
		fmt.Printf("%s is awesome!\n", language)
	}
	function2("Golang")
}
