package main

import "fmt"

var string1 = "I'm a VARIABLE outside of main() function!"

const string2 = "I'm a CONSTANT outside of main() function!"

func main() {

	num1, num2 := 10, 12

	fmt.Println(string1)
	fmt.Println(string2)

	sum := num1 + num2

	fmt.Println("Sum between 'num1' and 'num2':", sum)

}
