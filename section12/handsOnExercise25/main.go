package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var randNum int = rand.Intn(250)
	fmt.Println("The variable 'randNum' has value", randNum)

	switch {
	case randNum <= 100:
		fmt.Println("Between 0 and 100")
	case randNum >= 101 && randNum <= 200:
		fmt.Println("Between 101 and 200")
	case randNum >= 201:
		fmt.Println("Between 201 and 250")
	}

}
