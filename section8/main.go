package main

import (
	"fmt"
	"mymodule/section8/visibilityingo"
)

func main() {
	fmt.Println("Do you know how many Earths can be fit inside the Sun?")

	visibilityingo.SizeOfSun() // SizeOfSun() can be imported inside here because it starts with a capital letter!
	// This might be weird at first sight, but if you try to import the variable 'sunSize' directly from 'variableNonVisible.go', you could see that
	// it is not possible

	// This WOULD NOT work:
	/*
		visibilityingo.sunSize
	*/

	// But, inside 'variableVisible.go' there is a variable called 'SaturnSize'. It starts with a capital letter, so you can import it!
	// Like this:

	saturnSize := visibilityingo.SaturnSize
	fmt.Println("\nFun fact: you could fit", saturnSize, "Earth-sizes spheres inside the planet Saturn!")

	// Or just use it directly inside the 'fmt.Println()' function, like this:
	/*
		fmt.Println("\nFun fact: you could fit", visibilityingo.SaturnSize, "Earth-sizes spheres inside the planet Saturn!")
	*/
}
