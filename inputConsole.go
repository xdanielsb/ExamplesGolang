//Examples input from console
package main

import "fmt"

func main() {

	inputValues()

}

func inputValues() {

	fmt.Println("Please enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println("The double number is: ", output)

}
