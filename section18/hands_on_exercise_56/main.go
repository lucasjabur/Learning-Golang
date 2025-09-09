package main

import "fmt"

func main() {

	person1 := struct {
		firstName string
		friends   map[string]int
		favDrinks []string
	}{
		firstName: "Lucas",
		friends: map[string]int{
			"Fábio":   25,
			"Luan":    24,
			"Manuela": 21,
		},
		favDrinks: []string{
			"water",
			"coffee with milk",
			"orange juice",
		},
	}

	fmt.Println(person1.firstName, ":")
	for key, value := range person1.friends {
		fmt.Printf("   %s: %d\n", key, value)
	}
	for index, value := range person1.favDrinks {
		fmt.Printf("   %d. %s\n", index, value)
	}

}
