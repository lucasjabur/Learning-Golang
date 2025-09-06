package main

import (
	"fmt"
	"math/rand"
)

func main() {

	numberOfIterations := 0
	for i := 0; i < 100; i++ {
		if randNum := rand.Intn(5); randNum == 3 {
			fmt.Printf("%d. randNum is 3\n", i)
			numberOfIterations++
		}
	}

	fmt.Println("Total number of times where 'randNum' was equal to 3:", numberOfIterations)
}
