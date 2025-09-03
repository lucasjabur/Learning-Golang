package main

import "fmt"

type Person struct {
	first_name            string
	last_name             string
	fav_ice_cream_flavors []string
}

func main() {
	person1 := Person{
		first_name:            "Lucas",
		last_name:             "Jabur",
		fav_ice_cream_flavors: []string{"chocolate", "lime", "vanilla"},
	}

	person2 := Person{
		first_name:            "Manuela",
		last_name:             "Alvarenga",
		fav_ice_cream_flavors: []string{"strawberry", "chocolate", "pineapple"},
	}

	// fmt.Printf("Name: %s %s\nFavorite ice cream flavors:\n", person1.first_name, person1.last_name)
	// for index, value := range person1.fav_ice_cream_flavors {
	// 	fmt.Printf("   %d. %s\n", index+1, value)
	// }

	// fmt.Printf("\nName: %s %s\nFavorite ice cream flavors:\n", person2.first_name, person2.last_name)
	// for index, value := range person2.fav_ice_cream_flavors {
	// 	fmt.Printf("   %d. %s\n", index+1, value)
	// }

	// There are two different ways of creating and populating a map structure:
	// 1°: the way I did
	/*

		people_map := make(map[string]Person)
		people_map[person1.last_name] = person1
		people_map[person2.last_name] = person2

	*/

	// 2°: the way the professor did
	people_map := map[string]Person{
		person1.last_name: person1,
		person2.last_name: person2,
	}

	// The professor's solution for the main part of the exercise:
	/*

		for _, value := range people_map {
			fmt.Println(value)
			for _, flavors := range value.fav_ice_cream_flavors {
				fmt.Println(value.first_name, value.last_name, flavors)
			}
		}

	*/

	// My solution for the main part of the exercise is presented below.
	// I tried to create something different, and ended up doing this:
	/*

		flavors := person1.fav_ice_cream_flavors
		counter := 0

		for key, value := range people_map {
			fmt.Printf("%s: %s\n", key, value)
			for index, value := range flavors {
				fmt.Printf("   %d. %s\n", index, value)
				counter++
				if counter == len(flavors) {
					flavors = person2.fav_ice_cream_flavors
					break
				}
			}
		}

	*/

	// But that is not the best way to code this solution.
	// Here is an improvement of my solution using something I saw on the professor's solution:

	for key, value := range people_map {
		fmt.Printf("%s: %s\n", key, value)
		for index, flavors := range value.fav_ice_cream_flavors {
			fmt.Printf("   %d. %s\n", index, flavors)
		}
	}
}
