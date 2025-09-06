package main

import "fmt"

func main() {
	slice1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	slice2 := slice1[:5]
	fmt.Printf("%#v\n", slice2)

	slice3 := slice1[5:]
	fmt.Printf("%#v\n", slice3)

	slice4 := slice1[2:7]
	fmt.Printf("%#v\n", slice4)

	slice5 := slice1[1:6]
	fmt.Printf("%#v\n", slice5)

	fmt.Println("--------------------------")

	// Or just:
	fmt.Printf("%#v\n", slice1[:5])
	fmt.Printf("%#v\n", slice1[5:])
	fmt.Printf("%#v\n", slice1[2:7])
	fmt.Printf("%#v\n", slice1[1:6])
}
