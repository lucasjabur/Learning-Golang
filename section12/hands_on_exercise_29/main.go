package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// 1°:

	for i := 0; i < 100; i++ {
		fmt.Printf("%d \t", i)
	}
	fmt.Println("")

	// 2°:

	controlVariable := 0

	for controlVariable < 100 {
		randNum1 := rand.Intn(10)
		randNum2 := rand.Intn(10)

		switch {
		case randNum1 < 4 && randNum2 < 4:
			fmt.Printf("%d. Both are less than 4\n", controlVariable+1)
		case randNum1 > 6 && randNum2 > 6:
			fmt.Printf("%d. Both are greater than 6\n", controlVariable+1)
		case randNum1 >= 4 && randNum1 <= 6:
			fmt.Printf("%d. randNum1 is between 4 and 6\n", controlVariable+1)
		case randNum2 != 5:
			fmt.Printf("%d. randNum2 is not equal to 5\n", controlVariable+1)
		default:
			fmt.Printf("%d. None of the previous cases were met!\n", controlVariable+1)
		}
		controlVariable++
	}
}
