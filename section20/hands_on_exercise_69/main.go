package main

import "fmt"

var function1 = func(value int) {
	for i := 0; i < value; i++ {
		fmt.Println(i)
	}
}

func main() {

	function1(5)

}
