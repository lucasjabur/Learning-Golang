package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	age        int
}

type student struct {
	person
	university         string
	course             string
	graduation_student bool
}

func main() {

	/*

		person1 := person{
			first_name: "Lucas",
			last_name:  "Jabur",
			age:        23,
		}

		person2 := person{
			first_name: "Manuela",
			last_name:  "Alvarenga",
			age:        22,
		}

		fmt.Printf("Person 1: %#v, of type %T\n", person1, person1)
		fmt.Printf("Person 2: %#v, of type %T\n", person2, person2)

		fmt.Printf("%s %s is %d years old!\n", person1.first_name, person1.last_name, person1.age)
		fmt.Printf("%s %s is %d years old!\n", person2.first_name, person2.last_name, person2.age)

	*/

	// This is called 'embedded struct'.
	// We have a struct inside another struct ('person' it is inside 'student')
	student1 := student{
		person: person{
			first_name: "Lucas",
			last_name:  "Jabur",
			age:        23,
		},
		university:         "Federal University of Uberlandia (UFU)",
		course:             "Computer Science",
		graduation_student: true,
	}

	person2 := person{
		first_name: "Manuela",
		last_name:  "Alvarenga",
		age:        22,
	}

	fmt.Printf("Student 1: %#v, of type %T\n", student1, student1)
	fmt.Printf("Person 2: %#v, of type %T\n", person2, person2)

	fmt.Println(student1.person, student1.university, student1.course, student1.graduation_student)
	fmt.Println(student1.person.first_name, student1.university, student1.course, student1.graduation_student)

	// You can build an 'anonymous struct' as well, instead of declaring a struct, for example:
	// You basically substitute 'person' by its definition:
	person1 := struct {
		first_name string
		last_name  string
		age        int
	}{ // This is called a 'composite literal'
		first_name: "Sophia",
		last_name:  "Cardoso",
		age:        21,
	}

	fmt.Printf("\nPerson 1: %#v, of type %T\n", person1, person1)

}
