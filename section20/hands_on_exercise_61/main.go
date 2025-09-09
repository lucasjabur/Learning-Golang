package main

import "fmt"

type person struct {
	firstName string
	age       int
}

func (person person) speak() {
	fmt.Printf("Hi! My name is %s and I'm %d!\n", person.firstName, person.age)
}

func main() {

	person1 := person{
		firstName: "Lucas",
		age:       23,
	}

	person1.speak()

}
