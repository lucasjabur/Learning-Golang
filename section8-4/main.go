package main

import puppy "github.com/lucasjabur/Puppy"

// So, now we're able to work with tags
// But, as you can see in the 'go.mod' file, the imported version of 'Puppy' repo is v0.0.0
// It means that we're not importing the most recent version available, because 'Puppy' is at version v1.4.0 already.

// Let's see if we can update it by using the commands we already know:
/*
	1°: 'go mod tidy' --> nothing has changed!
	2°: 'go get github.com/lucasjabur/Puppy@v1.4.0' --> it worked! Now, we're getting v1.4.0 of 'Puppy' repo

*/

// Now, let's test it!

func main() {
	puppy.From14()
}
