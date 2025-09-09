package main

import "fmt"

func incrementer() func() int { // the function 'incrementer()' is closing over the other function. This is called 'closure'

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

	// What's expected is that the output would be: '1, 1, 1, 1, 1'
	// But, what actually happens is: '1, 2, 3, 4, 5'
	// What's going on? Simple...
	// When line 17 is executed, the command 'value := 0' is executed, then the anonymous inner function is created and capture that variable ('CLOSURE!!!!')
	// Then the anonymous function is returned by the main function ('incrementer()') and gets store in the variable 'function'
	// So, when 'function()' is called, what's I'm calling is actually the anonymous inner function (the return of function 'incrementer()'),
	// not the function 'incrementer()' itself!

}
