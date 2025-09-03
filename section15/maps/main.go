package main

import "fmt"

func main() {

	// The map declaration syntax follows this pattern: 'map[key type]value type'

	map1 := map[string]int{
		"Lucas":   23,
		"Manuela": 22,
		"Isabela": 15,
	}

	fmt.Println("The age of Lisa is:", map1["Lisa"])
	fmt.Printf("%#v\n", map1)

	map2 := make(map[string]int)
	map2["Fabiana"] = 53
	map2["André"] = 58
	map2["Sophia"] = 21

	fmt.Printf("%#v\n", map2)
	fmt.Printf("Length: %d\n", len(map2))

	fmt.Println("--------------------")

	// Something interesting is that Golang orders alphabetically by default!

	// Now, let's see how to range over a map:

	// Ranging over the map and printing the keys and the values stored in it:
	for key, value := range map2 {
		fmt.Printf("%s: %v\n", key, value)
	}
	fmt.Println("--------------------")

	// Ranging over the map and printing just the values:
	for _, value := range map2 {
		fmt.Printf("%v\n", value)
	}
	fmt.Println("--------------------")

	// Ranging over the map and printing just the keys:
	for key := range map2 {
		fmt.Printf("%v\n", key)
	}
	fmt.Println("--------------------")

	fmt.Printf("Antes de deletar: %#v\n", map2)
	// Now, how to delete an element from a map?
	delete(map2, "Sophia")
	fmt.Printf("Depois de deletar: %#v\n", map2)
	fmt.Println("--------------------")

	// Something that can confuse some people about maps in Golang, is that if you do not type the key name correctly, it won't give an error, like so:
	fmt.Println("Acessing keys that don't exist: ")
	delete(map2, "Andrey") // this will return '0', and that's because the value stored in 'map2' is of type 'int'.
	// If it was a 'boolean' it would return 'false', if it was a 'string' it would return "", and so on...
	fmt.Println(map2["Fabian"])
	fmt.Println("--------------------")

	// So, how would you know if 'Andrey' exists and was just born or not?
	// Well, using the 'comma, ok' idiom! Like the following:

	value, ok := map2["Andrey"]
	if ok {
		fmt.Printf("The value is: %v\n", value)
	} else {
		fmt.Println("The key don't exist!") // this will be the output!
	}

	// It also can be written with the 'statement, statement' idiom:
	if value, ok := map2["Andrey"]; ok {
		fmt.Printf("The value is: %v\n", value)
	} else {
		fmt.Println("The key don't exist!")
	}
	fmt.Println("--------------------")

	// It is important to remeber: the importance of 'statement, statement' idiom is the scope limitation of the variables 'value' and 'ok'!
	// This can be seen directly on the examples above, where I used the same variable names in two different scenarios.
	// This is cool, because it helps the developer mantain the code cleaner

	// Something interesting is that you can use the '++' operator on maps as well:
	map3 := make(map[string]int)
	fmt.Println(map3["Lucas"]) // output here: '0'
	map3["Lucas"]++
	map3["Lucas"]++
	map3["Lucas"]++
	map3["Lucas"]++
	fmt.Println(map3["Lucas"]) // output here: 4

	// NOTE FOR THE FUTURE: TRY TO DO THE FOLLOWING EXERCISE:
	// -> Read the 'romeo-and-juliet.txt' file
	// -> Get all the words from it and treat them as necessary (removing "-", "'", "...", etc from the words)
	// -> Display how many times each word appear in the book
	// -> Display what is the word that appears the most

}
