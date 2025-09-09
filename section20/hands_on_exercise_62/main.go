package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	length int
	width  int
}

type circle struct {
	radius float64
}

func (rectangle rectangle) areaCalc() float64 {
	area := rectangle.length * rectangle.width
	return float64(area)
}

func (circle circle) areaCalc() float64 {
	area := math.Pi * math.Pow(circle.radius, 2)
	return area
}

type shape interface {
	areaCalc() float64 // need to specify the type of the return here!
}

func info(shape shape) {
	shape.areaCalc()
}

func main() {

	rectangle1 := rectangle{
		length: 5,
		width:  5,
	}

	circle1 := circle{
		radius: 2,
	}

	rectangleArea := rectangle1.areaCalc()
	fmt.Println(rectangleArea)
	circleArea := circle1.areaCalc()
	fmt.Println(circleArea)

	info(rectangle1)
	info(circle1)

}
