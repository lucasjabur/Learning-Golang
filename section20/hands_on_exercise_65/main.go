package main

import "fmt"

func factorial(value int) int {
	if value == 0 {
		return 1
	}
	return value * factorial(value-1)
}

func main() {

	value := 6
	factorial := factorial(value)
	fmt.Printf("Factorial of %d is: %d\n", value, factorial)

}
