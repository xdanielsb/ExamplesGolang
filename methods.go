//Methods Go

package main

import "fmt"

func main() {
	c1 := Complex{real: 1, imaginary: 5}
	c1.show()

	p1 := Person{name: "Daniel", id: 845}
	p1.greet()
}

type Complex struct {
	real      float64
	imaginary float64
}

//This is the receiver
func (c Complex) show() {
	fmt.Printf("%f , %fi \n", c.real, c.imaginary)
}

type Person struct {
	name string
	id   int
}

func (p Person) greet() {
	fmt.Printf("Hi, I am a person\n")
}
