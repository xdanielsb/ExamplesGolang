//Sentence defer in go
package main

import "fmt"

func main() {
	//The important aspect is declare two sentences together
	//but the sentence that have defer is going te execute at the
	//end of the current function
	defer task2()
	task1()
}

func task1() {
	fmt.Println("Task 1")
}

func task2() {
	fmt.Println("Task 2")
}
