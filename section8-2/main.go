package main

import (
	"fmt"

	puppy "github.com/lucasjabur/Puppy"
)

func main() {
	string1 := puppy.Bark()
	string2 := puppy.Barks()

	// If you try to the new functionalities that includes the Dog repo's funtion now, it won't work:

	/*
		string3 := puppy --> the only options available here is 'Bark()' and 'Barks()'
	*/
	// That's because the commit that is being referenced is an old one that includes just the two inicial functions
	// So, just do: 'go get github.com/lucasjabur/Puppy@8fefe76', where '8fefe76' is the number of the new commit, in the root directory

	// Now, Dog is also being imported, because Puppy uses Dog as a direct dependency!

	string3 := puppy.BigBark()
	string4 := puppy.BigBarks()

	fmt.Println(string1)
	fmt.Println(string2)

	fmt.Println(string3)
	fmt.Println(string4)
}
