package main

import "fmt"

func main() {

	// Hello, Gopher!

	fmt.Println("Hello, Gopher!")

	// Formatted print

	const name, age = "Lucas", 23
	fmt.Printf("%s is %d years old!\n", name, age)

	fmt.Printf("\tAnd the type is %T and %T\n", name, age)

	// Raw string literals

	fmt.Println(`
		This is what's called
		'Raw string literals'.
		They are a sequence of characters between
		back quotes \n.
		As you can see, the '\n' does not work here!
	`)

	// To print variables along side with text, you can use fmt.Println() as well, like:

	const name2, age2 = "Manuela", 21
	fmt.Println(name2, "is", age2, "years old!")
}
