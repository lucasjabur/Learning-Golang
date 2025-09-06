package main

import "fmt"

type person struct {
	first_name string
}

type student struct {
	person
	course string
}

func (person person) introduceYourself() {
	fmt.Println("I am,", person.first_name)
}

func (student student) introduceYourself() {
	fmt.Println("I am,", student.first_name) // 'student' are also people because I defined it inside the 'Student' struct
}

type human interface { // an 'interface' is a set of method signatures
	introduceYourself()
}

// So, ANY value that has the method 'introduceYourself()', that is inside the 'human' interface, WILL ALSO BE OF TYPE 'human'.
// In Go, values can be of different types simultaneously
// AND THAT IS: POLYMORPHISM!

func saySomething(human human) {
	human.introduceYourself()
}

func main() {

	student1 := student{
		person: person{
			first_name: "Lucas",
		},

		course: "Computer Science",
	}

	person2 := person{
		first_name: "Manuela",
	}

	// person1.introduceYourself()
	student1.introduceYourself()
	person2.introduceYourself()

	saySomething(student1)
	saySomething(person2)

}
