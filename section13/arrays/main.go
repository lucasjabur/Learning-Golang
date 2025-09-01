package main

import "fmt"

func main() {
	array1 := [3]int{42, 43, 44}
	fmt.Println(array1)

	array2 := [...]string{"Hello,", "Gophers!"}
	fmt.Println(array2)

	var array3 [2]int
	array3[0] = 10
	array3[1] = 12
	fmt.Println(array3)

	fmt.Printf("%T ", array1)
	fmt.Printf("%T ", array2)
	fmt.Printf("%T\n", array3)
}
