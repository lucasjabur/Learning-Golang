package main

import "fmt"

func main() {

	person1 := struct {
		first_name string
		friends    map[string]int
		fav_drinks []string
	}{
		first_name: "Lucas",
		friends: map[string]int{
			"Fábio":   25,
			"Luan":    24,
			"Manuela": 21,
		},
		fav_drinks: []string{
			"water",
			"coffee with milk",
			"orange juice",
		},
	}

	fmt.Println(person1.first_name, ":")
	for key, value := range person1.friends {
		fmt.Printf("   %s: %d\n", key, value)
	}
	for index, value := range person1.fav_drinks {
		fmt.Printf("   %d. %s\n", index, value)
	}

}
