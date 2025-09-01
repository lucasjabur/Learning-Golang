package main

import (
	"fmt"
	"sort"
)

func main() {

	slice1 := []string{"Hello,", "World!"}
	fmt.Println(slice1)

	fmt.Println("\nLooping over a slice with for range loop:")
	// This WON'T work:
	/*

		for value := range slice1 {
			fmt.Printf("%s\n", value)
		}

	*/

	// You need to have the 'index' variable alongside with the 'value' variable to print out the slice:
	for index, value := range slice1 {
		fmt.Printf("%d. %s\n", index+1, value)
	}

	// But, what if I don't want to print out the index variable? How do I do that?
	// Simple: use the 'blank identifier'
	for _, value := range slice1 {
		fmt.Printf("%s\n", value)
	}

	// The for range loop operation returns the index value, so it cannot be removed.
	// You need to use the 'blank identifier' to ignore it

	fmt.Println("\nPrinting by index position:")
	fmt.Println(slice1[0])
	fmt.Println(slice1[1])

	fmt.Println("\nLooping over a slice with 'default' for loop:")
	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}

	// Now, how do we put more elements inside a slice?
	fmt.Println("\nAppending values to a slice:")
	slice2 := []int{42, 43, 44}
	fmt.Println(slice2)
	slice2 = append(slice2, 45)
	fmt.Println(slice2)

	// And how to slice a slice?
	fmt.Println("\nSlicing a slice:")
	slice3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("slice3: %#v\n", slice3)
	fmt.Println("---------------------------------------------")

	// [inclusive:exclusive]
	fmt.Printf("slice3: %#v\n", slice3[2:8]) // value stored on index 0 is printed out, but the value of index 4 is not
	fmt.Println("---------------------------------------------")

	// [:exclusive]
	fmt.Printf("slice3: %#v\n", slice3[:6]) // it goes from the beggining 'till the indicated index, but the value of that index is not shown
	fmt.Println("---------------------------------------------")

	// [inclusive:]
	fmt.Printf("slice3: %#v\n", slice3[3:]) // it goes from the value of the indicated index (including it) 'till the end of the slice
	fmt.Println("---------------------------------------------")

	// [:]
	fmt.Printf("slice3: %#v\n", slice3[:]) // gives the entire slice
	fmt.Println("---------------------------------------------")

	// How about deleting from a slice?
	fmt.Println("\nDeleting from a slice:")
	slice4 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Using append() function:
	slice4 = append(slice4[:4], slice4[5:]...)
	// WHAT'S HAPPENING HERE?

	// We are appending 'slice4[5:]...' to 'slice4[:4]'

	// 'slice4[5:]' gives us a slice from index 5 until the end of the slice4, so: '{5, 6, 7, 8, 9}'
	// but '{5, 6, 7, 8, 9}' is a set of values, not the values itself
	// so we have to tell the compiler that we want those values, like this: 'slice4[5:]...'
	// Now, we have the values to append: '5, 6, 7, 8, 9'.

	// The sintax 'append(slice4[:4], slice4[5:]...)' means that we want to append those values to the slice 'slice4[:4]'
	// But what is 'slice4[:4]'? Since '[:4]' includes the slice from the beggining without including the index 4 value, it gives us:
	// '{0, 1, 2, 3}', a slice with the values from indexes 0 to 3 of slice4 (excluding index 4, as we wanted)

	// Now, let's complete the append operation:
	// We are appending '5, 6, 7, 8, 9' to the slice '{0, 1, 2, 3}', and that gives us: '{0, 1, 2, 3, 5, 6, 7, 8, 9}'
	// That's the same output as if we just removed the value of index 4, and that's exactly what we wanted!

	fmt.Printf("slice4: %#v\n", slice4)
	fmt.Println("---------------------------------------------")

	// Now, another way of creating a slice:

	slice5 := make([]int, 0, 5)
	// The 'make()' function receives 2 arguments by default:
	// 		- the type of the elements that are gonna be stored in the slice
	//      - the size of the slice
	// But by passing those 3 arguments: 'make([]int, 0, 5)', what we are saying is:
	// I'm not storing any values in the slice yet, but I want it to be able to store 5 elements
	// So the underlying array (because a slice is built on top of an array) of the slice has 5 places of memory allocated to store 5 values of type int.

	fmt.Println("\nCreating a slice with the make() function:")
	fmt.Printf("slice5: %#v\n", slice5)
	fmt.Printf("Size of slice5: %v\n", len(slice5))     // 0, because there is no values stored in it
	fmt.Printf("Capacity of slice5: %v\n", cap(slice5)) // 5, because it can hold 5 elements

	slice5 = append(slice5, 0, 1, 2, 3, 4)
	fmt.Println("---------------------------------------------")
	fmt.Printf("slice5: %#v\n", slice5)
	fmt.Printf("Size of slice5: %v\n", len(slice5))
	fmt.Printf("Capacity of slice5: %v\n", cap(slice5))

	// What happen if we try to store more values than we defined? An error?
	slice5 = append(slice5, 5, 6, 7)
	fmt.Println("---------------------------------------------")
	fmt.Printf("slice5: %#v\n", slice5)
	fmt.Printf("Size of slice5: %v\n", len(slice5))
	fmt.Printf("Capacity of slice5: %v\n", cap(slice5))

	// It works, but the size and the capacity changes!
	// But, how does it change? Does it follows a pattern?
	// ... YES!
	// Let's take a look:

	for i := 0; i < 40; i++ {
		slice5 = append(slice5, i)

		if len(slice5) == cap(slice5) {
			fmt.Println("---------------------------------------------")
			fmt.Printf("slice5: %#v\n", slice5)
			fmt.Printf("Size of slice5: %v\n", len(slice5))
			fmt.Printf("Capacity of slice5: %v\n", cap(slice5))
		}
	}

	fmt.Println("---------------------------------------------")
	fmt.Printf("slice5: %#v\n", slice5)
	fmt.Printf("Size of slice5: %v\n", len(slice5))
	fmt.Printf("Capacity of slice5: %v\n", cap(slice5))

	// Here we can see that: when the size of elements stored in the array gets bigger than the capacity,
	// the capacity doubles in value!
	// Capacity growth: 5 -> 10 -> 20 -> 40 -> 80 ...

	// Now, let's learn about multidimensional slices:
	// What are they? Simply a slice of slices. So instead of storing values, it stores another slice!

	fmt.Println("\nCreating a multidimensional slice from existing slices:")
	slice6 := []string{"Lucas", "Taekwondo", "Development", "Reading"}
	slice7 := []string{"Manuela", "Reading", "Drawing", "Biotech"}
	slice8 := [][]string{slice6, slice7}
	fmt.Printf("slice8: %#v\n", slice8)

	fmt.Println("\nCreating a multidimensional slice by hand:")
	slice9 := [][]int{{0, 1, 2}, {3, 4, 5}}
	fmt.Printf("slice8: %#v\n", slice9)

	fmt.Println("\nUnderstanding how slices work under the table:")
	// Imagine the following situation:
	slice10 := []int{0, 1, 2, 3, 4, 5}
	slice11 := slice10

	fmt.Printf("%#v\n", slice10)
	fmt.Printf("%#v\n", slice11)
	fmt.Println("---------------------------------------------")

	// Here, it is obvious that slice10 and slice11 are going to be the same
	// We are saying it explicitly that slice11 = slice10

	// But why do the following happen?
	slice10[0] = 42
	fmt.Printf("%#v\n", slice10)
	fmt.Printf("%#v\n", slice11)
	fmt.Println("---------------------------------------------")

	// That's why a slice has a underlying array, and that means that a slice has a pointer pointing to an array
	// So, 'slice10' is pointing to an underlying array, let's call it 'underlying_array10', for example
	// In other words: 'slice10' ---> 'underlying_array10'
	// When we say: 'slice11 := slice10', what the Golang compiler sees is: 'slice11' ---> 'underlying_array10'
	// The result of that: every modification we do to 'slice10' will also modify 'underlying_array10' and, since
	// 'underlying_array10' is the underlying array for 'slice11' as well, it will also modify 'slice11'.

	// Now, let's play a little more:
	fmt.Println("\nUnderstanding how slices work under the table (PART 2):")
	slice12 := []int{0, 1, 2, 3, 4, 5}
	slice13 := make([]int, 6)
	copy(slice13, slice12) // (destination, source)

	fmt.Printf("%#v\n", slice12)
	fmt.Printf("%#v\n", slice13)
	fmt.Println("---------------------------------------------")

	slice12[0] = 42
	fmt.Printf("%#v\n", slice12)
	fmt.Printf("%#v\n", slice13)
	fmt.Println("---------------------------------------------")

	// Now, 'slice13' has a different underlying array from 'slice12'
	// We copied the values from 'slice12' and pasted them inside 'slice13'.
	// 'slice13' does not reference 'slice12'.

	fmt.Println("\nUnderstanding how slices work under the table (PART 3 - USING FUNCTIONS):")
	slice14 := []float64{0.0, 3.0, 2.0, 1.0}
	fmt.Println("Median:", sliceMedianWithReference(slice14)) // after being sorted, median must be: (1.0 + 2.0) = 1.5
	fmt.Println("Slice now:", slice14)

	slice15 := []float64{0.0, 3.0, 2.0, 1.0, 5.0}
	fmt.Println("\nMedian:", sliceMedianWithoutReference(slice15)) // after being sorted, median must be: 1.0
	fmt.Println("Slice now:", slice15)
}

// In this function, 'slice' and 'slice14' are pointing to the same underlying array
func sliceMedianWithReference(slice []float64) float64 {
	sort.Float64s(slice)
	index := len(slice) / 2
	if len(slice)%2 == 1 {
		return slice[index/2]
	}
	return (slice[index-1] + slice[index]) / 2
}

// And here, we are creating a copy of 'slice15', so 'slice16' and 'slice15' do not point to the same underlying array
func sliceMedianWithoutReference(slice []float64) float64 {
	slice16 := make([]float64, len(slice))
	copy(slice16, slice)

	sort.Float64s(slice16)
	index := len(slice16) / 2
	if len(slice16)%2 == 1 {
		return slice16[index/2]
	}
	return (slice16[index-1] + slice16[index]) / 2
}
