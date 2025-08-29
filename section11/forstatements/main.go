package main

import "fmt"

func main() {
	// FOR STATEMENT

	num1, num2 := 42, 5
	fmt.Printf("num1 = %v and num2 = %v\n", num1, num2)

	for i := 0; i < 5; i++ {
		fmt.Printf("Counting to five: %v \t first for loop\n", i)
	}

	for num2 < 10 {
		fmt.Printf("num2 is %v \t\t second for loop\n", num2)
		num2++
	}

	for {
		fmt.Printf("num2 is %v \t\t third for loop\n", num2)
		if num2 > 20 {
			break // step out of the loop
		}
		num2++
	}

	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue // jumps over to the next step of the loop ignoring the rest of the code
		}
		fmt.Println("Counting even numbers: ", i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("--")
		for j := 0; j < 5; j++ {
			fmt.Printf("Outer loop %v \t inner loop %v\n", i, j)
		}
	}

	slice1 := []int{10, 20, 30, 40, 50, 60}
	for i, v := range slice1 {
		fmt.Println("Ranging over a slice:", i, v) // 'i' gives the position of the value 'v' in the slice
	}

	map1 := map[string]int{
		"Lucas":   23,
		"Manuela": 22,
	}
	for k, v := range map1 {
		fmt.Println("Ranging over a map:", k, v) // 'k' gives the key and 'v' the value stored in the map
	}
}
