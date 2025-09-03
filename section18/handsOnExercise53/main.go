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

	fmt.Printf("Name: %s %s\nFavorite ice cream flavors:\n", person1.first_name, person1.last_name)
	for index, value := range person1.fav_ice_cream_flavors {
		fmt.Printf("   %d. %s\n", index+1, value)
	}

	fmt.Printf("\nName: %s %s\nFavorite ice cream flavors:\n", person2.first_name, person2.last_name)
	for index, value := range person2.fav_ice_cream_flavors {
		fmt.Printf("   %d. %s\n", index+1, value)
	}

}
