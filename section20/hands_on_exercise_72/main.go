package main

import (
	"fmt"
	"math"
)

func powinator() func() float64 {
	value := 0.0
	return func() float64 {
		value++
		return math.Pow(value, 2.0)
	}
}

func higherPower() func() float64 {
	value := 2.0

	return func() float64 {
		value = math.Pow(value, 2.0)
		return value
	}
}

func main() {
	function1 := higherPower()
	fmt.Println(function1())
	fmt.Println(function1())
	fmt.Println(function1())
	fmt.Println(function1())

	function2 := powinator()
	fmt.Println(function2())
	fmt.Println(function2())
	fmt.Println(function2())
	fmt.Println(function2())

}

// This was made on the lesson about 'closure':

/*

func incrementer() func() int {
	value := 0

	return func() int {
		value++
		return value
	}
}

func main() {
	function := incrementer()
	fmt.Println(function())
	fmt.Println(function())
	fmt.Println(function())
	fmt.Println(function())
	fmt.Println(function())
	fmt.Println(function())
}

*/
