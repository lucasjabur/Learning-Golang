package main

import "fmt"

type Person struct {
	first_name string
}

func (person Person) speak() {
	fmt.Println("I am,", person.first_name)
}

func main() {

	// Basic syntax of functions in Golang: 'func (receiver) identifier(parameters) (returns) { code }'
	// Important note: everything in Go is PASSED BY VALUE!

	function_example1()
	function_example2("I'm a parameter")
	function_example3("Parameter 1", "Parameter 2")
	string1, string2 := function_example4("STRING1", "STRING2")
	fmt.Printf("string1 now: %s\n", string1)
	fmt.Printf("string2 now: %s\n", string2)

	// You can pass a slice as parameter to a function that receives a variadic parameter as well.
	// For example, let's use the function 'sum()':

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := sum(values...) // this is called 'unfurling a slice'

	/*

		sum := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
		fmt.Println("The sum is:", sum)

	*/

	fmt.Println("The sum is:", sum)

	sentence := bringWordsTogether("Golang", "is", "the", "best", "programming language", "of all time!")
	fmt.Println(sentence)

	// Now, let's learn about the 'defer' statement

	function_example1()
	defer function_example2("Hello!")

	function_example3("Good bye!", "Golang!")

	// The output must've been: output of 'function_example1', output of 'function_example2' and output of 'function_example3'
	/*

		I'm a function and I don't have any parameters or returns!
		I'm a function and I have one parameter 'Hello!', but no returns!
		I'm a function and I have two parameter 'Good bye!' and 'Golang!' (I could have more by the way...), but no returns!

	*/
	// But, the output that we got was: output of 'function_example1', output of 'function_example3' and output of 'function_example2'
	/*

		I'm a function and I don't have any parameters or returns!
		I'm a function and I have two parameter 'Good bye!' and 'Golang!' (I could have more by the way...), but no returns!
		I'm a function and I have one parameter 'Hello!', but no returns!

	*/
	// And that because of the 'defer' command!
	// This command delays the execution of the function associated with it, causing it to be executed only at the end of the function where it is being called!
	// In this case, 'function_example2' is inside main(), so 'function_example2' will only be executed at the end of 'main()'.
	// The 'defer' statement is useful for closing a file, instead of adding the command to close the file at the end of the code, you can do it right after openning it
	// with the 'defer' statement before.
	// It helps with the clarity of the code and prevents you from forgetting to close it afterwards.

	// Now, let's understand how to build methods in Golang:

	person1 := Person{
		first_name: "Lucas",
	}

	person2 := Person{
		first_name: "Manuela",
	}

	person1.speak()
	person2.speak()

	// 'speak()' is a method and it has a 'person Person' as a 'receiver'.
	// It means that every 'Person' variable will have access to the 'speak()' function (method)

}

// Some examples of functions:

func function_example1() {
	fmt.Println("I'm a function and I don't have any parameters or returns!")
}

func function_example2(string1 string) {
	fmt.Printf("I'm a function and I have one parameter '%s', but no returns!\n", string1)
}

func function_example3(string1, string2 string) {
	fmt.Printf("I'm a function and I have two parameter '%s' and '%s' (I could have more by the way...), but no returns!\n", string1, string2)
}

func function_example4(string1, string2 string) (string, string) {
	fmt.Println("I'm a function and I have two parameter and also two returns! Look!")
	temp := string1
	string1 = string2
	string2 = temp

	return string1, string2
}

// You can build a function that receives an unlimited amount of values!
// That is called 'variadic parameter'

func sum(elements ...int) int { // this variadic parameter MUST BE the FINAL INCOMING PARAMETER of the function
	// You could have: 'math(string1 string, operation rune, elements ...int)', for example
	fmt.Println(elements)        // that is a slice of 'int' values
	fmt.Printf("%T\n", elements) // as you can see here

	sum := 0
	for _, value := range elements {
		sum += value
	}

	return sum
}

func bringWordsTogether(words ...string) string { // this functions is just an usage example of variadic parameters of another type

	var sentence string
	for _, value := range words {
		temp := value + " "
		sentence += temp

	}

	return sentence
}
