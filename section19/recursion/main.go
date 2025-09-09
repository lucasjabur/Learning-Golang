package main

import "fmt"

func factorial(value int) int {
	fmt.Println("This is the value: ", value)
	if value == 0 {
		return 1
	}

	return value * factorial(value-1)
}

func main() {

	factorial := factorial(4)
	fmt.Println("The factorial of 4 is:", factorial)

}
