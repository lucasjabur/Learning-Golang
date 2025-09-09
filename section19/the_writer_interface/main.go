package main

import (
	"bytes"
	"fmt"
)

// type person struct {
// 	firstName string
// }

// func (person person) writeOut(writer io.Writer) error {
// 	_, err := writer.Write([]byte(person.firstName))

// 	return err
// }

func main() {
	// file, err := os.Create("output.txt")
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }
	// defer file.Close()

	// slice1 := []byte("Hello, Gophers!") // you can go from a string to a slice of bytes, and from a slice of bytes to a string
	// _, err = file.Write(slice1)
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }

	buffer1 := bytes.NewBufferString("Hello ")
	fmt.Println(buffer1.String()) // a value of type 'buffer' implements the function 'String()', so it also implements the interface 'Stringer'
	buffer1.WriteString("Gophers!")
	fmt.Println(buffer1.String())
	buffer1.Reset()
	buffer1.WriteString("It's Monday") // it also implements the function 'WriteString()', for example...
	fmt.Println(buffer1.String())

	buffer1.Write([]byte("... let's work!"))

}
