//Functions
package main

import "fmt"

func main() {
	grades := []float64{98, 93, 77, 82, 83}
	result := average(grades)
	fmt.Println(result)
	result = float64(suma(1, 2))
	fmt.Println(result)
	fmt.Println(getNames())
	fmt.Println(sumValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(grades...)
}

//First form to declare a function
func average(grades []float64) float64 {
	total := 0.0
	for _, value := range grades {
		total += value
	}
	return (total / float64(len(grades)))

}

//Another form to return something
func suma(numero1 int, numero2 int) (total int) {
	total = numero1 + numero2
	return
}

//Returning multiples values

func getNames() (name string, surname string) {
	name = "Daniel"
	surname = "Santos"
	return
}

//Declaring a function where it does not know how many parameters receive
//... helps me to do that
func sumValues(grades ...int) (total int) {
	total = 0
	for _, values := range grades {
		total += values
	}
	return
}
