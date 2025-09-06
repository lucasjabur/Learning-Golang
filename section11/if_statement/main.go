package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// IF STATEMENT

	num1, num2 := 42, 5
	fmt.Printf("num1 = %v and num2 = %v\n", num1, num2)

	if num1 < 42 {
		fmt.Println("num1 is less than the meaning of life!")
	} else if num1 > 42 {
		fmt.Println("num1 is greater than the meaning of life!")
	} else {
		fmt.Println("num1 is equal to the meaning of life!")
	}

	// This is called 'statement; statement' idiom:

	if num3 := 2 * rand.Intn(num1); num3 >= num1 {
		fmt.Printf("num3 is %v and that is GREATER THAN or EQUALS num1 which is %v\n", num3, num1)
	} else {
		fmt.Printf("num3 is %v and that is LESS THAN num1 which is %v\n", num3, num1)
	}
	// Here we're declaring the variable 'num3' on the definition of the 'if statement' and applying a condition for it to run

	// This won't run:
	/*
		fmt.Println(num3)
	*/
	// That's because 'num3' is declared in the 'if statement', it means that the scope of 'num3' is the 'if statement' scope

	// This is called 'comma, ok' idiom:

	students := map[string]float64{
		"Lucas":   8.8,
		"Manuela": 10.0,
	}

	grade, ok := students["Manuela"]
	fmt.Println(grade, ok)

	grade, ok = students["Isabela"]
	fmt.Println(grade, ok)
	// In this case, it is being used to show if the student was found or not in the 'students' map (similar to Python's 'dictionary' type)
}
