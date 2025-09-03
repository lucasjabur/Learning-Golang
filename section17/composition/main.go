package main

import "fmt"

type Engine struct {
	power     float64
	torque    float64
	fuel_type string
}

type Car struct {
	Engine
	brand string
	model string
	color string
	year  int
}

func (e *Engine) Start(engine Engine) {
	fmt.Printf("The engine has %.2f horses of power, %.2f kgf*m of torque and uses %s as fuel.\n", engine.power, engine.torque, engine.fuel_type)
}

func CarInfo(car Car) {
	fmt.Printf("The car is a %d %s %s %s. What a car!\n", car.year, car.color, car.brand, car.model)
}

func main() {
	car1 := Car{
		brand: "Ford",
		model: "Fusion",
		color: "black",
		year:  2022,
	}
	engine1 := Engine{
		power:     248.0,
		torque:    39.0,
		fuel_type: "flex",
	}
	CarInfo(car1)
	car1.Start(engine1) // calling the Start() method from the embedded 'Engine' struct (inside 'Car' struct)
}

// In Golang, there are no classes, you create 'types' instead.
// In Golang, also, you don't instantiate, you create a value of a 'type'
// But other object orianted functionalities can be used here:
//    -> Encapsulation
//    -> Inheritance (embedded types)
//    -> Polymorphism (interfaces)
//    -> Overriding
