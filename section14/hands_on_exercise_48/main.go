package main

import "fmt"

func main() {

	// I tried to make it without creating other slices like the professor did, so this is my solution:

	slice1 := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "I'm 008"}}
	fmt.Printf("%#v\n", slice1)

	for index, value := range slice1 {
		fmt.Printf("%d. %v\n", index, value)
		for i := 0; i < len(value); i++ {
			fmt.Printf("    %d. %v\n", i, value[i])
		}
	}

	// Professor's solution:

	// slice2 := []string{"James", "Bond", "Shaken, not stirred"}
	// slice3 := []string{"Miss", "Moneypenny", "I'm 008"}

	// slice4 := [][]string{slice2, slice3}
	// fmt.Printf("%#v\n", slice4)

	// for index, value := range slice4 {
	// 	fmt.Printf("%d. %v\n", index, value)
	// 	for index2, value2 := range value {
	// 		fmt.Printf("    %d. %v\n", index2, value2)
	// 	}
	// }

	// Overall, the output is the same!
	// The difference is that my solution is more memory efficient.

}
