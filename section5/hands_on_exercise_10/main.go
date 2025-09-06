/*

Write a program that uses the following:
- var for zero value
- short declaration operator
- multiple initializations
- var when specificity is required
- blank identifier

Print items as necessary to make the program interesting

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	varForZeroValue()
	shortDeclarationOperator()
	multipleInitializations()
	varWithSpecificity()

	var sqrt9, sqrt25 float64
	_, sqrt9, sqrt25 = blankIdentifier() // blank identifier used for ignoring the first return value of the function
	fmt.Printf("\nsqrt9 = %.2f and sqrt25 = %.2f\n", sqrt9, sqrt25)
}

func varForZeroValue() {
	var int1, int2 int
	var float1 float64
	var string1 string
	var rune1 rune
	var bool1 bool

	fmt.Printf("int1 = %d and int2 = %d\n", int1, int2)
	fmt.Printf("float1 = %f\n", float1)
	fmt.Printf("string1 = %s\n", string1)
	fmt.Printf("rune1 = %c\n", rune1)
	fmt.Printf("bool1 = %v\n", bool1)
}

func shortDeclarationOperator() {
	num1, num2 := 10, 20
	sum := num1 + num2

	fmt.Printf("\nIf num1 = %d and num2 = %d, then sum = %d\n", num1, num2, sum)
}

func multipleInitializations() {
	var num3, num4, num5 = 5, 10, 15

	fmt.Printf("\nnum3 = %d, num4 = %d and num5 = %d\n", num3, num4, num5)
}

func varWithSpecificity() {
	var string2 string = "Golang is awesome!"

	fmt.Println("\nEveryone know that", string2)
}

func blankIdentifier() (float64, float64, float64) {
	num6, num7, num8 := 4.0, 9.0, 25.0
	return math.Sqrt(num6), math.Sqrt(num7), math.Sqrt(num8)
}
