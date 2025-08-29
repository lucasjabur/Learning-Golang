package main

import "fmt"

func main() {
	slice1 := []int{42, 43, 44, 45, 46, 47}
	for i, v := range slice1 {
		fmt.Printf("Ranging over a slice: %v, %v\n", i, v)
	}
}
