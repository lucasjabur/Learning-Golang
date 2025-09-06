package main

import "fmt"

func main() {
	slice1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	slice1 = append(slice1, 52)
	fmt.Printf("%#v\n", slice1)

	slice1 = append(slice1, 53, 54, 55)
	fmt.Printf("%#v\n", slice1)

	slice2 := []int{56, 57, 58, 59, 60}

	// This will work! It was the solution I came up with to solve the problem of appending 'slice2' to 'slice1, but...

	/*
		for index := range slice2 {
			slice1 = append(slice1, slice2[index])
		}
		fmt.Printf("%#v\n", slice1)
	*/

	// There is a better way!

	// To append 'slice2' to 'slice1', you should do like the following:
	slice1 = append(slice1, slice2...)
	fmt.Printf("%#v\n", slice1)

}
