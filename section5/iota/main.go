package main

import "fmt"

const (
	_ = iota // this first value using 'iota' is '0', and the rest below is incremented by 1
	a        // so this is 1
	b        // this is 2
	c        // this is 3 and so on...
	d
	e
	f
)

func main() {
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a)
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b)
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c)
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d)
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e)
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f)

	anotherIotaExample()
}

// From Efective Go
type ByteSize float64 // here a new type it is being created (named type) 'ByteSize' and its underlying type is the underlying types of 'float64'
// this means that the underlying type is recursively traced back until a predeclared type or a type literal is reached
// In this case, 'float64' is already a predeclared type, so the underlying type of 'ByteSize' is 'float64'

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // iota here is '1' = 2^10
	MB                             // iota here is '2': 1 << (10 * 2) = 2^20
	GB                             // iota here is '3': 1 << (10 * 3)  = 2^30 and so on...
	TB
	PB
	EB
	ZB
	YB
)

func anotherIotaExample() {
	fmt.Printf("\nKB: %.2f \t \t \t \tBinary: %b\n", KB, KB)
	fmt.Printf("MB: %.2f \t \t \t \tBinary: %b\n", MB, MB)
	fmt.Printf("GB: %.2f \t \t \tBinary: %b\n", GB, GB)
	fmt.Printf("TB: %.2f \t \t \tBinary: %b\n", TB, TB)
	fmt.Printf("PB: %.2f \t \tBinary: %b\n", PB, PB)
	fmt.Printf("EB: %.2f \t \tBinary: %b\n", EB, EB)
	fmt.Printf("ZB: %.2f \t \tBinary: %b\n", ZB, ZB)
	fmt.Printf("YB: %.2f\tBinary: %b\n", YB, YB)
}
