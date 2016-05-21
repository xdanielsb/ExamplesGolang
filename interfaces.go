//Interfaces helps me to explicit the similiraties between structures
//Roughly Like an interface in java

package main

import (
	"fmt"
	"math"
)

func main() {
	c1 := Circle{r: 1.0}
	c2 := Circle{r: 2.0}
	r1 := Rectangle{l1: 1.0, l2: 2.0}
	r2 := Rectangle{l1: 5.0, l2: 3.0}
	area(c1, c2, r1, r2)
}

type Shape interface {
	area() float64
}

type Circle struct {
	r float64
}

type Rectangle struct {
	l1, l2 float64
}

func (c Circle) area() (area float64) {
	area = c.r * c.r * math.Pi
	return
}

func (r Rectangle) area() (area float64) {
	area = r.l1 * r.l2
	return
}

func area(shapes ...Shape) float64 {
	var area float64
	for _, shape := range shapes {
		area += shape.area()
		fmt.Printf("%f \n", shape.area())
	}
	return area
}
