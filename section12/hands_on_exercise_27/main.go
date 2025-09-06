package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randNum1 := rand.Intn(10)
	randNum2 := rand.Intn(10)

	fmt.Printf("randNum1 = %d and randNum2 = %d\n", randNum1, randNum2)

	if randNum1 < 4 && randNum2 < 4 {
		fmt.Println("Both are less than 4")
	} else if randNum1 > 6 && randNum2 > 6 {
		fmt.Println("Both are greater than 6")
	} else if randNum1 >= 4 && randNum1 <= 6 {
		fmt.Println("randNum1 is between 4 and 6")
	} else if randNum2 != 5 {
		fmt.Println("randNum2 is not equal to 5")
	} else {
		fmt.Println("None of the previous cases were met!")
	}

}
