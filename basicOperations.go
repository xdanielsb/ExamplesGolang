//Operations with diferent values"

package main

import "fmt"

func main() {
	fmt.Println("Hello this is me the winner of the house")
	operationsStrings()
	operationsLogic()
	operationNumbers()
}

func operationsStrings() {
	//this the lenStrings of the string
	fmt.Println(len("Ana de Dios es una ganadora"))
	//get an specific character of the string
	fmt.Println("Ana de Dios es una ganadora"[7])
	//Concatenate two strings
	fmt.Println("Ana de Dios " + "es una ganadora")
}

func operationsLogic() {

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}

func operationNumbers() {
	fmt.Println(849498988 * 54654556516516165415.0)
	fmt.Println(84.0 / 9822646486.0)

}
