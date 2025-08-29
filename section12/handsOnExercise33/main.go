package main

import "fmt"

func main() {

	for variable := 0; variable < 20; variable++ {
		if variable%2 != 0 {
			continue
		}
		fmt.Println(variable)
	}
}
