package main

import "fmt"

func main() {
	map1 := make(map[string][]string)
	map1[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	map1[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	map1[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}

	map1[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	fmt.Println("Before deleting: ")
	for key, value := range map1 {
		fmt.Printf("%s: %v\n", key, value)
	}

	delete(map1, `moneypenny_jenny`)

	fmt.Println("\nAfter deleting: ")
	for key, value := range map1 {
		fmt.Printf("%s: %v\n", key, value)
	}
}
