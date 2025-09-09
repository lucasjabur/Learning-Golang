package main

import "fmt"

func function1() int {
	return 42
}

func function2() func() int {
	return func() int {
		return 43
	}
}

func function3() func() int {
	return function1 // you could do this as well!
}

func main() {

	return1 := function1()
	fmt.Println(return1)          // type 'int'
	fmt.Printf("%T\n", function1) // type 'func() int'

	return2 := function2()
	fmt.Println(return2())
	fmt.Printf("%T\n", function2) // type 'func() func() int'
	fmt.Printf("%T\n", return2)   // type func() int'

	return3 := function3()
	fmt.Println(return3())
	fmt.Printf("%T\n", function3) // type 'func() func() int'
	fmt.Printf("%T\n", return3)   // type 'func() int'
}
