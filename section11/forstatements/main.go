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
}
