//Arrays and other things
package main

import "fmt"

func main() {
	arrays()
	castingAmongVariables()
	goThroughArray()
	anotherWay()
}

//Casting Variables
func castingAmongVariables() {
	var number int = 845
	number2 := float64(number)
	fmt.Println(number2)
}

//Declaring arrays
func arrays() {
	var values [5]int
	values[4] = 54
	fmt.Println(values)
}

//Another way to initialize an array
func anotherWay() {
	array := [6]float64{1, 2, 3, 4, 5, 6} //The last comma is required
	fmt.Println("The array initialize of a fancy form is", array)
}

//Go through an array
func goThroughArray() {
	var array [5]float64
	var total float64 = 0

	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5
	for _, value := range array {
		total += value
	}
	fmt.Println("The sum of the array is ", total)

}
