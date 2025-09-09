package main

import (
	"fmt"
	"log"
	"os"
)

func read_file(file_name string) ([]byte, error) { // 'read_file()' is a wrapper function
	file, err := os.ReadFile(file_name) // because it envolves the original "read file" function
	if err != nil {
		return nil, fmt.Errorf("error in read_file() func: %s", err)
	}
	return file, nil
}

// A wrapper function adds functionality to antoher function. In this case, it treats the error that 'os.ReadFile()' returns
// This helps the programmer to understand exactly what happened, for example, using and manipulating the error message

func main() {
	file, err := read_file("poem.txt")
	if err != nil {
		log.Fatalf("read_file in main(): %s", err)
	}

	fmt.Println(file)
	fmt.Println(string(file))
}
