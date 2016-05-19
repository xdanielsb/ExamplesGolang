//THis code show how to use the codes public and recover
package main

import "fmt"

func main() {
	example()
}

func example() {

	//It is necesary to declare of this form because panic end the program
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("Error in a fuck thing")
}
