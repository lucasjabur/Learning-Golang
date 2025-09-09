package main

import "fmt"

func sum(value1 int, value2 int) int {
	return value1 + value2
}

func returningIntAndString(year int, name string) (int, string) {
	age := 2025 - year

	introducingMyself := "Hi, I'm " + name

	return age, introducingMyself
}

func main() {
	function1 := sum(10, 20)
	fmt.Printf("Sum: %d\n", function1)

	return1, return2 := returningIntAndString(2002, "Lucas")
	fmt.Println("First return (int):", return1)
	fmt.Println("Second return (string):", return2)
}
