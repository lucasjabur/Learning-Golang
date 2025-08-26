package section2

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
}
