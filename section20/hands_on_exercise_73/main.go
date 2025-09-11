package main

import (
	"fmt"
	"time"
)

func timeFunc(function func()) {
	startTimer := time.Now()
	function()
	timeElapsed := time.Since(startTimer)
	fmt.Println("\n", timeElapsed)
}

func doSomething() {
	for i := 0; i < 10_000; i++ {
		fmt.Println(i)
	}
}

func main() {
	timeFunc(doSomething)
}
