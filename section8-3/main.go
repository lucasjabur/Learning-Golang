package main

import "fmt"

func main() {
	fmt.Println("Hello! We're working on versions now!.")

	// If you run on the terminal: 'git tag', nothing will show up as response.
	// That's because we're not tagging our repo with versions
	// So, now we're going to learn how to do it:
	/*
		1. 'git add ...'
		2. 'git tag v1.0.0', for example
		3. 'git commit -m "..."'
		4. 'git push'
	*/
	// Now, 'git tag' must get 'v1.0.0' as output
}
