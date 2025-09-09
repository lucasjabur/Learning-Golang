package main

import "fmt"

func add(value1, value2 int) int {
	return value1 + value2
}

func subtract(value1, value2 int) int {
	return value1 - value2
}

func doMath(value1, value2 int, function func(int, int) int) int {
	return function(value1, value2)
}

func main() {
	// 'Callback' is passing a function as an argument to another function!

	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	sum := doMath(45, 12, add)
	fmt.Println(sum)

	sub := doMath(453, 213, subtract)
	fmt.Println(sub)

}
