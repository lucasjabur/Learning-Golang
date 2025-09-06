package main

import (
	"fmt"
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

func main() {

	book1 := book{
		title: "Red Rising",
	}

	var value count = 42

	fmt.Println(book1)
	fmt.Println(value)

}
