package main

func main() {
	// Functions:
	/*

		- Purpose of functions:
			- modular code
			- abstracting code
			- code reusability (DRY)

		- Parameters vs Arguments:
			- parameters: values showed on the signature of the function
			- arguments: values passed to the function

		- Syntax of a function:
			- 'func (r receiver) functionName(p parameters) (returns) { code }'

		- Variadic parameters:
			- variadic parameters are an infinite amount of values that can be passed to a function, like so:
				- 'func sum(elements ...int) int { code }'
				- remembering that 'elements' is a slice of 'int' values
				- so, you can do like:
					- 'sumReturn := sum(1, 2, 3, 4, 5)', or
					- 'values := []int{1, 2, 3, 4, 5}; sumReturn := sum(values...)'

		- Unfurling a slice:
			- 'values := []int{1, 2, 3, 4, 5}; sumReturn := sum(values...)' here I am unfurling a slice
			- the use of '...' expands the slice into individual arguments when calling a veriadic function

		- 'defer' statement:
			- the 'defer' statement delays the execution of a function or a method to execute at the end of the function where it is put in
				func main() {
					function1()
					defer function2()
					function3()
				}
				- the natural order of execution here would be: 'function1', 'function2' and 'function3'
				- but 'defer function2()' makes the 'function2' be the last one to be executed
				- and it will be executed at the end of the execution of 'main()' function

		- Methods:
			- take the following code as example:

				type person struct {
					firstName string
				}

				func (person person) speak() {
					fmt.Println("My name is ", person.firstName)
				}

				- in this case, 'speak()' is a method of the 'person' struct
				- a method is a function with the 'receiver' argument
				- the value specified in the 'receiver' means that it implements that method

		- Interfaces and polymorphism:
			- take the following code as example:

				type person struct {
					firstName string
				}

				type student struct {
					person
					course string
				}

				func (person person) introduceYourself() {
					fmt.Println("I am,", person.firstName)
				}

				func (student student) introduceYourself() {
					fmt.Println("I am,", student.firstName)
				}

				type human interface {
					introduceYourself()
				}

				func saySomething(human human) {
					human.introduceYourself()
				}

				- an 'interface' is populated by methods
				- these methods can be implemented by the values of that type ('interface')
				- when the function 'saySomething()' is called, the instruction 'human.introduceYourself()' will execute
					the 'introduceYourself()' method that is connected to the type of human passed initially, so:

					- 'saySomething(person1)' will print the name of the 'person1'
					- 'saySomething(student1)' will print the name of the 'student1'

				- it is important to say that every value that implements the method 'introduceYourself()' (that is inside the 'human' interface)
					will also be of type 'human', so 'student' and 'person' are also 'human's
					- this is polymorphism!

		- Anonymous function:
			- this is a function that doesn't have an identifier, neither a receiver:

				func main() {
					func(name string, year int) int {
						fmt.Println("My name is", name)
						age := 2025 - year

						return age
					}("Lucas", 2002)
				}

		- Function expressions:
			- we can assign a function to a variable:

				func main() {

					function1 := exampleFunction() {
						fmt.Println("This is an example!")
					}
					function1()
				}

		- Returning a function:
			- functions are first class citizens, just like variables and types, for example
			- so we can return a function just as we can return a variable of type 'int':

				func returningFunction() func() {
					return func() {
						fmt.Println("Hello! I'm a function and a return!")
					}
				}

		- Callbacks:
			- 'callbacks' are a function being passed as argument to another function:

				func add(value1 int, value2 int) int {
					return value1 + value2
				}

				func subtract(value1 int, value2 int) int {
					return value1 - value2
				}

				func doMath(value1 int, value2 int, function func(int, int) int) int {
					return function(value1, value2)
				}

				func main() {

					sum := doMath(10, 10, add)
					fmt.Println(sum)

					sub := doMath(10, 10, subtract)
					fmt.Println(sub)
				}

		- Closure:
			- 'closure' is one scope enclosing another, basically
			- a variable can be declared in the outer scope and be accessed in the inner scope
			- it helps us limit the scope of the variables

				func incrementer() func() int {
					value := 0
					return func() {
						value++
						return value
					}
				}

			- here the scope of the variable 'value' (outer scope) is enclosing the scope of the 'func()' (inner scope) that is being returned
			- that way, the variable is accessed inside the inner scope without any problems!

		- Recursion:
			- basically, 'recursion' is when a function calls itself

				func factorial(value int) int {
					if value == 0 {
						return 1
					}
					return value * factorial(value - 1)
				}

			- every recursive function needs a base case that will stop the recursion
			- this prevents an infinite recursive call and makes the function work as expected
	*/
}
