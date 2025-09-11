package main

import "fmt"

func main() {
	// Anonymous func that returns the name passed as argument
	func(name string) {
		fmt.Println("My name is", name)
	}("Lucas")
}
