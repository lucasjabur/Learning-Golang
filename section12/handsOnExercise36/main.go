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
}
