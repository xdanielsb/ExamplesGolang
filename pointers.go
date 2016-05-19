//This code explain how works the pointers in go
package main

import "fmt"

func main() {
	number := 3
	fmt.Println(number)
	passByValue(number)
	fmt.Println(number)
	passByReference(&number)
	fmt.Println(number)
}

func passByValue(number int) {
	number = 1
}

func passByReference(number *int) {
	*number = 9
}
