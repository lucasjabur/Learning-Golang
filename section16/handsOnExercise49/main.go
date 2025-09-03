package main

import "fmt"

func main() {
	map1 := make(map[string][]string)
	map1[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	map1[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	map1[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}

	for key, value := range map1 {
		fmt.Printf("%s: %v\n", key, value)
		for i := 0; i < len(value); i++ {
			fmt.Printf("    %d. %v\n", i, value[i])
		}
	}

}
