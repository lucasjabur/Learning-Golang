package main

import "fmt"

func main() {
	slice1 := []string{"Almond Biscotti Café", "Banana Pudding", "Balsamic Strawberry (GF)",
		"Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough",
		"Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)",
		"Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)",
		"Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)",
		"Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ",
		"Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different",
		"Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)",
		"Wolverine Tracks (GF)"}

	fmt.Println(slice1)
	fmt.Println("\nThe length of slice1 is:", len(slice1))

	fmt.Println("\nNow using a for range loop:")

	for index, value := range slice1 {
		fmt.Printf("%d. %s\n", index+1, value)
	}
}
