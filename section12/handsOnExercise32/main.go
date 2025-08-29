package main

import (
	"fmt"
)

func main() {

	variable := 5

	for {
		fmt.Printf("variable is %v\n", variable)
		if variable > 20 {
			break
		}
		variable++
	}
}
