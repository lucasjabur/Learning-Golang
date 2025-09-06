package main

import "fmt"

type Engine struct {
	electric bool
}

type Vehicle struct {
	engine Engine
	make   string
	model  string
	doors  int
	color  string
}

func main() {

	vehicle1 := Vehicle{
		engine: Engine{
			electric: true,
		},
		make:  "Toyota",
		model: "Corolla Cross",
		doors: 4,
		color: "white",
	}

	vehicle2 := Vehicle{
		engine: Engine{
			electric: false,
		},
		make:  "Mitsubishi",
		model: "Eclipse GT",
		doors: 2,
		color: "black",
	}

	fmt.Println(vehicle1)
	fmt.Println(vehicle2)

	fmt.Printf("\nThe %s %s it is %s and has %d doors\n", vehicle1.make, vehicle1.model, vehicle1.color, vehicle1.doors)
	fmt.Printf("The %s %s it is %s and has %d doors\n", vehicle2.make, vehicle2.model, vehicle2.color, vehicle2.doors)

}
