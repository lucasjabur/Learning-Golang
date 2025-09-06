package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// SWITCH STATEMENT:

	num1, num2 := 42, 5
	fmt.Printf("num1 = %v and num2 = %v\n", num1, num2)

	switch num1 {
	case 41:
		fmt.Println("num1 is less than the meaning of life!")
	case 42:
		fmt.Println("num1 is equal to the meaning of life!")
	case 43:
		fmt.Println("num1 is one algarism bigger than the meaning of life!")
	default:
		fmt.Println("This is the default case for num1!")
	}

	switch {
	case num1 == 41:
		fmt.Println("num1 is less than the meaning of life!")
	case num1 == 42:
		fmt.Println("num1 is equal to the meaning of life!")
	case num1 == 43:
		fmt.Println("num1 is one algarism bigger than the meaning of life!")
	default:
		fmt.Println("This is the default case for num1!")
	}

	// SELECT STATEMENT:

	channel1 := make(chan int)
	channel2 := make(chan int)

	duration1 := time.Duration(rand.Int63n(250))
	duration2 := time.Duration(rand.Int63n(250))

	go func() {
		time.Sleep(duration1 * time.Millisecond)
		channel1 <- 41
	}()

	go func() {
		time.Sleep(duration2 * time.Millisecond)
		channel2 <- 42
	}()

	// A 'select' statement let us handle multiple channel operations concurrently.
	// It allows a Go program to wait for communication on multiple channels simultaneously and execute the code
	// associated with the first channel operation that becomes ready.

	select {
	case value1 := <-channel1:
		fmt.Println("Value from channel 1", value1)
	case value2 := <-channel2:
		fmt.Println("Value from channel 2", value2)
	}
}
