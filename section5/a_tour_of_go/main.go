package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// This will work:
var bool4, bool5, bool6 bool

// This WON'T work:
// bool7 := true

func main() {

	// Playing with some functionalities:

	fmt.Println(rand.Intn(51654654)) // returns a pseudo-random number from '0' to 'n-1': [0,n)
	fmt.Println(math.Sqrt(4))        // returns the square root of 'n'
	fmt.Println(math.Pi)             // returns the constant Pi

	fmt.Println(add(10, 12))
	sayHello()

	string1, string2 := swap("Hello", "Golang")
	fmt.Println(string1, string2)

	fmt.Println(split(42))

	// Another cool feature of Go is that you can use 'var' statement to declare a list of variables:

	var bool1, bool2, bool3 bool
	fmt.Println(bool1, bool2, bool3) // the output will be 'false false false', because the variables didn't have values assigned to them
	fmt.Println(bool4, bool5, bool6) // you can declare those variables on 'package scope' (outside functions)
	// it is important to mention that the short declaration operator WILL NOT WORK OUTSIDE OF FUNCTIONS

	// When using 'var' statement you can include initializers to assign values to the variables
	// If an initializer is present, the type can be omitted, so:

	// This:
	var num1, num2 int = 1, 2 // here you are declaring the type of the variables
	// Is the same of:
	var num3, num4 = 1, 2 // here you are telling the compiler to guess what is the type of the variables
	// This is called type inference

	fmt.Printf("num1 = %d e num2 = %d\n", num1, num2)
	fmt.Printf("num3 = %d e num4 = %d\n", num3, num4)

	// Some other basic types of Golang:
	someBasicTypes()

	// A little more about convertion:
	convertingALittleBit()

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

// Now, let's play a little bit with functions!

// The structure here is kid of self-explanatory:

func add(x, y int) int { // you have to say what type are the function's parameters (int) as well as the type of the return (int)
	return x + y
}

func sayHello() {
	fmt.Println("Hello!")
}

func swap(string1, string2 string) (string, string) {
	return string2, string1
}

// Another way to return values is using named return values.
// Instead of declaring just the type of the return, you declare what variables will be returned by that function, like:

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // here you don't need to write the variables you want to return, the compiler aready knows
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func someBasicTypes() {
	fmt.Printf("Type: %T; Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T; Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T; Value: %v\n", z, z)
}

func convertingALittleBit() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Printf("Type: %T; Value: %v\n", x, x)
	fmt.Printf("Type: %T; Value: %v\n", y, y)
	fmt.Printf("Type: %T; Value: %v\n", z, z)
}

// An important concept in Golang: an untyped constant takes the type needed by its context

const (
	Big   = 1 << 100  // the binary number that is 1 followed by 100 zeros
	Small = Big >> 99 // shift it right again 99 places, so it's the same as 1 << 1, or 2
)

func needInt(num5 int) int {
	return num5*10 + 1
}
func needFloat(num5 float64) float64 {
	return num5 * 0.1
}
