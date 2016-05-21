//Methods Go

package main

import "fmt"

type Complex struct {
	real      float64
	imaginary float64
}

func main() {
	c1 := Complex{real: 1, imaginary: 5}
	c1.show()
}

func (c Complex) show() {
	fmt.Printf("%f , %fi", c.real, c.imaginary)
}
