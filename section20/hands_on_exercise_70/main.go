package main

import "fmt"

func outerFunc() func() int {
	return func() int {
		return 42
	}
}

func main() {

	function := outerFunc()
	fmt.Println(function())

}
