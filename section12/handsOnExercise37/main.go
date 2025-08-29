package main

import "fmt"

func main() {
	map1 := map[string]int{
		"Lucas":   23,
		"Manuela": 22,
	}
	for k, v := range map1 {
		fmt.Printf("%v has %v years old\n", k, v)
	}
	map2 := map1["Lucas"]
	fmt.Println(map2)

	map3 := map1["Q"]
	fmt.Println(map3)

	// It can be done like this:
	age1, ok1 := map1["Q"]
	fmt.Println(age1, ok1)

	// Or like this:
	if value, ok := map1["Q"]; !ok {
		fmt.Println("There is no Q, and here is the zero value of an int:", value)
	}

}
