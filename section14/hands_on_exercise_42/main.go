package main

import "fmt"

func main() {
	array1 := [5]int{0, 1, 2, 3, 4}

	for _, value := range array1 {
		fmt.Printf("%v, of type %T\n", value, array1)
	}
}
