package main

import (
	"fmt"
	"log"
	"os"
)

func readFile(fileName string) ([]byte, error) { // 'readFile()' is a wrapper function
	file, err := os.ReadFile(fileName) // because it envolves the original "read file" function
	if err != nil {
		return nil, fmt.Errorf("error in readFile() func: %s", err)
	}
	return file, nil
}

// A wrapper function adds functionality to antoher function. In this case, it treats the error that 'os.ReadFile()' returns
// This helps the programmer to understand exactly what happened, for example, using and manipulating the error message

func main() {
	file, err := readFile("poem.txt")
	if err != nil {
		log.Fatalf("readFile in main(): %s", err)
	}

	fmt.Println(file)
	fmt.Println(string(file))
}
