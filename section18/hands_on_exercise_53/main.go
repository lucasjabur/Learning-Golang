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

	fmt.Printf("Name: %s %s\nFavorite ice cream flavors:\n", person1.firstName, person1.lastName)
	for index, value := range person1.favIceCreamFlavors {
		fmt.Printf("   %d. %s\n", index+1, value)
	}

	fmt.Printf("\nName: %s %s\nFavorite ice cream flavors:\n", person2.firstName, person2.lastName)
	for index, value := range person2.favIceCreamFlavors {
		fmt.Printf("   %d. %s\n", index+1, value)
	}

}
