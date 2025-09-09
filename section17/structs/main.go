package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type student struct {
	person
	university        string
	course            string
	graduationStudent bool
}

func main() {

	/*

		person1 := person{
			firstName: "Lucas",
			lastName:  "Jabur",
			age:        23,
		}

		person2 := person{
			firstName: "Manuela",
			lastName:  "Alvarenga",
			age:        22,
		}

		fmt.Printf("Person 1: %#v, of type %T\n", person1, person1)
		fmt.Printf("Person 2: %#v, of type %T\n", person2, person2)

		fmt.Printf("%s %s is %d years old!\n", person1.firstName, person1.lastName, person1.age)
		fmt.Printf("%s %s is %d years old!\n", person2.firstName, person2.lastName, person2.age)

	*/

	// This is called 'embedded struct'.
	// We have a struct inside another struct ('person' it is inside 'student')
	student1 := student{
		person: person{
			firstName: "Lucas",
			lastName:  "Jabur",
			age:       23,
		},
		university:        "Federal University of Uberlandia (UFU)",
		course:            "Computer Science",
		graduationStudent: true,
	}

	person2 := person{
		firstName: "Manuela",
		lastName:  "Alvarenga",
		age:       22,
	}

	fmt.Printf("Student 1: %#v, of type %T\n", student1, student1)
	fmt.Printf("Person 2: %#v, of type %T\n", person2, person2)

	fmt.Println(student1.person, student1.university, student1.course, student1.graduationStudent)
	fmt.Println(student1.person.firstName, student1.university, student1.course, student1.graduationStudent)

	// You can build an 'anonymous struct' as well, instead of declaring a struct, for example:
	// You basically substitute 'person' by its definition:
	person1 := struct {
		firstName string
		lastName  string
		age       int
	}{ // This is called a 'composite literal'
		firstName: "Sophia",
		lastName:  "Cardoso",
		age:       21,
	}

	fmt.Printf("\nPerson 1: %#v, of type %T\n", person1, person1)

}
