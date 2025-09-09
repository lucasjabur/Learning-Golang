package main

import "fmt"

func sumVariadicParameter(elements ...int) int {
	sum := 0
	for _, value := range elements {
		sum += value
	}
	return sum
}

func normalSum(slice []int) int {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}

func main() {

	slice1 := []int{1, 2, 3, 4, 5}
	variadicSum := sumVariadicParameter(slice1...)
	fmt.Println("Sum of elements:", variadicSum)

	slice2 := []int{10, 20, 30, 40, 50}
	normalSum := normalSum(slice2)
	fmt.Println("Sum of elements:", normalSum)

}
