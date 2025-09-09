package main

import "fmt"

type Person struct {
	firstName          string
	lastName           string
	favIceCreamFlavors []string
}

func main() {
	person1 := Person{
		firstName:          "Lucas",
		lastName:           "Jabur",
		favIceCreamFlavors: []string{"chocolate", "lime", "vanilla"},
	}

	person2 := Person{
		firstName:          "Manuela",
		lastName:           "Alvarenga",
		favIceCreamFlavors: []string{"strawberry", "chocolate", "pineapple"},
	}

	// fmt.Printf("Name: %s %s\nFavorite ice cream flavors:\n", person1.firstName, person1.lastName)
	// for index, value := range person1.favIceCreamFlavors {
	// 	fmt.Printf("   %d. %s\n", index+1, value)
	// }

	// fmt.Printf("\nName: %s %s\nFavorite ice cream flavors:\n", person2.firstName, person2.lastName)
	// for index, value := range person2.favIceCreamFlavors {
	// 	fmt.Printf("   %d. %s\n", index+1, value)
	// }

	// There are two different ways of creating and populating a map structure:
	// 1°: the way I did
	/*

		peopleMap := make(map[string]Person)
		peopleMap[person1.lastName] = person1
		peopleMap[person2.lastName] = person2

	*/

	// 2°: the way the professor did
	peopleMap := map[string]Person{
		person1.lastName: person1,
		person2.lastName: person2,
	}

	// The professor's solution for the main part of the exercise:
	/*

		for _, value := range peopleMap {
			fmt.Println(value)
			for _, flavors := range value.favIceCreamFlavors {
				fmt.Println(value.firstName, value.lastName, flavors)
			}
		}

	*/

	// My solution for the main part of the exercise is presented below.
	// I tried to create something different, and ended up doing this:
	/*

		flavors := person1.favIceCreamFlavors
		counter := 0

		for key, value := range peopleMap {
			fmt.Printf("%s: %s\n", key, value)
			for index, value := range flavors {
				fmt.Printf("   %d. %s\n", index, value)
				counter++
				if counter == len(flavors) {
					flavors = person2.favIceCreamFlavors
					break
				}
			}
		}

	*/

	// But that is not the best way to code this solution.
	// Here is an improvement of my solution using something I saw on the professor's solution:

	for key, value := range peopleMap {
		fmt.Printf("%s: %s\n", key, value)
		for index, flavors := range value.favIceCreamFlavors {
			fmt.Printf("   %d. %s\n", index, flavors)
		}
	}
}
