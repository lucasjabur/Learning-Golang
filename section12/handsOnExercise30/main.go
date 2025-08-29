package main

import (
	"fmt"
	"math/rand"
)

func randNumOperations() {
	randNum := rand.Intn(5)

	switch randNum {
	case 0:
		fmt.Printf("\nrandNum = %d\n", randNum)
	case 1:
		fmt.Printf("\nrandNum = %d\n", randNum)
	case 2:
		fmt.Printf("\nrandNum = %d\n", randNum)
	case 3:
		fmt.Printf("\nrandNum = %d\n", randNum)
	default:
		fmt.Printf("\nrandNum = %d\n", randNum)
	}
}

func main() {

	for i := 0; i < 42; i++ {
		randNumOperations()
		fmt.Println("Iteration: ", i+1)
	}

}
