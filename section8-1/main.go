package main

// Using "go get https://github.com/lucasjabur/Puppy@latest" to get the latest commit made on the Puppy repository
/*

Output:
	go: downloading github.com/lucasjabur/Puppy v0.0.0-20250829111349-b73440164613
	go: added github.com/lucasjabur/Puppy v0.0.0-20250829111349-b73440164613

*/

import (
	"fmt"

	puppy "github.com/lucasjabur/Puppy"
)

func main() {
	string1 := puppy.Bark()
	string2 := puppy.Barks()

	fmt.Println(string1)
	fmt.Println(string2)
}

// After importing and using the import, Golang puts a comment in 'go.mod' saying that the imported package is 'indirect'
// It means that the import may or may not be used inside this Go project.
// It is not an error, but to correct it you can run 'go mod tidy', AFTER USING THE IMPORT
// So it becomes a 'direct' dependency, meaning that it is a dependecy that is directly used by the project
