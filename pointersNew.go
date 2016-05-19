//Operator new pointers

package main

import "fmt"

//Another way to get a pointer
func main() {
	number := new(int)
	fmt.Println(*number)
	newOperator(number)
	fmt.Println(*number)
}

func newOperator(number *int) {
	*number = 5
}
