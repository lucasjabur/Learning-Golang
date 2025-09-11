package main

import "fmt"

func Square(value int) int {
	return value * value
}

func printSquare(function func(int) int, value int) string {

	square := function(value)
	return fmt.Sprintf("The square value of %d is: %v\n", value, square)
}

func main() {

	fmt.Println(printSquare(Square, 9))

}
