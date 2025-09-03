package main

import "fmt"

func main() {
	slice1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, value := range slice1 {
		fmt.Println(value)
		fmt.Printf("%v, of type %T\n", value, value)
	}
}
