package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	firstName string
}

func (person person) writeOut(writer io.Writer) error {
	_, err := writer.Write([]byte(person.firstName))

	return err
}

func main() {

	person1 := person{
		firstName: "Lucas",
	}

	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer file.Close()

	var buffer bytes.Buffer
	person1.writeOut(file)
	person1.writeOut(&buffer)    // the bytes buffer for 'writeOut()' to work need to be a pointer to a buffer, so you need to pass the memory address of the buffer
	fmt.Println(buffer.String()) // prints the buffer content
}
