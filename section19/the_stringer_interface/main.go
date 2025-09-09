package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

type count int

// Remembering: the 'receiver' of a function attaches that function to the 'type' expressed inside as the 'receiver'
// In this case below, the function is attached to the 'book' struct

func (book book) String() string {
	return fmt.Sprint("The title of the book is: ", book.title)
}

func (count count) String() string {
	return fmt.Sprint("This is the number: ", strconv.Itoa(int(count))) // here I am converting the 'count' value to an 'int' and then conveting it to a 'string' using the ASCII Table
}

func logInfo(str fmt.Stringer) { // ANYTHING THAT IS IMPLEMENTING THE 'fmt.Stringer' interface CAN BE PASSED IN HERE!
	log.Println("LOG FROM SECTION19:", str.String())
}

func main() {

	book1 := book{
		title: "Red Rising",
	}

	var value count = 42

	fmt.Println(book1)
	fmt.Println(value)

	log.Println(book1) // format that prints the date and the time that the message was printed
	log.Println(value)

	logInfo(book1) // here I am expanding the functionality of the function 'String()'
	logInfo(value) // that is known as a 'wrapper function', or just 'wrapper'
	// It is a function that provides an additional layer of abstraction or functionality around an existing function or method.
	// In this case, when the 'logInfo' function receives a parameter of type 'book' as 'str', the 'log.Println()' command uses the 'String()' function
	// that is attributed to the 'book' type. The same happens when a parameter of type 'count' is passed.

}
