package main

import "fmt"

func main() {
	// Declaring variables
	// Golang is a statically typed language, so:

	var num1 int = 10
	fmt.Println(num1)
	num1 = 15
	fmt.Println(num1)

	// But, you can use the short declaration operator so you don't have to declare it's type, like:

	num2 := 34.6
	fmt.Printf("%f\n", num2)
	fmt.Printf("With less zeros: %.2f or %.4f, for example\n", num2, num2)
	fmt.Printf("%f it's a %T number\n", num2, num2)

	// Golang is smart, it won't let you declare a variable and not use it. This was made to avoid code pollution!
	// This WON'T work at all:

	/*

		num3, num4, num5 := 20, 30, 40
		fmt.Printf("num3 = %d e num4 = %d", num3, num4)

	*/

	// But this will:

	num3, num4, _, string1 := 20, 30, 40, "Golang"
	fmt.Printf("num3 = %d e num4 = %d, string1 = %s\n", num3, num4, string1)

	// Using '_' is important so you won't have to use, for example, a return value from a function that you don't
	// need in a certain part of your code.

	// Another cool feature about Golang:

	var num6 int // This will assign the integer value of 0 (zero) to the variable 'num6':
	fmt.Printf("num6 = %d\n", num6)

	var num7 float64 // This will assign the float64 value of 0 (zero) to the variable 'num7':
	fmt.Printf("num6 = %.2f\n", num7)

	// It works with all types, boolean (false), int (0), float (0.000000), string (""), function, pointers, interfaces, etc...
	// All of the last ones are "nil"

	// You can manipulate the numbers as well
	// You can print the hexadecimal type of number, or it's binary correspondent, such as:

	_, num9 := 42, 13
	fmt.Printf("num8 is %d in the decimal form\n", num9)
	fmt.Printf("num8 is %b in the binary form\n", num9)
	fmt.Printf("num8 is %x in the hexadecimal form\n", num9)
	fmt.Printf("num8 is %#X in the hexadecimal form\n", num9)

	// There are other forms that you can use to play with the numbers like %e and %E for scientific notation,
	// %X for upper-case hexadecimal numbers and %#X for 0X... hexadecimal format

	// String and bytes can also be played with:

	string2 := "Awesome"
	fmt.Printf("string2 is %s in the normal form\n", string2)
	fmt.Printf("string2 is %q in the double quoted form\n", string2)
	fmt.Printf("string2 is %x in the base16 lower-case form\n", string2)

	// Other types have that feature as well!

	// Let's imagine the following scenario: I have a variable of type float32 and want it to store a float64 numerical value. Can it be done?
	// It depends! The way of doing it shown below WON'T work at all!

	/*

		var num10 float32 = 23.564
		var num11 float64

		num11 = num10
		fmt.Printf("num11 = %.4f\n", num11)

	*/

	// To work, you'll have to convert it into a float64 variable, like this:

	var num10 float32 = 23.564
	var num11 float64

	num11 = float64(num10)
	fmt.Printf("num10 = num11 = %.4f\n", num11)
}
