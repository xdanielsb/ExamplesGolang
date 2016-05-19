//Variables and control flow statements
package main

import "fmt"

func main() {
	classicForm()
	alternativeForm()
	keyPoints()
	multipleVariables()
}

//Classic form of Declarating variables
func classicForm() {
	//Declaration of the String
	var name string = "Hello, I am a variable"
	fmt.Println(name)

	//Anther way to assign the value of a string
	var surname string
	surname = "My last name is Fonseca"
	fmt.Println(surname)
	fmt.Println(name == surname)
}

//Another form to Declarate variables
func alternativeForm() {
	//With this form is not necesary declarate the type of the variables
	name := "I am a modern variable"
	fmt.Println(name)
	//Third form
	var surname = "I am contenporaneous variable"
	fmt.Println("This is the text: ", surname)
}

//Creating constants in the program
func keyPoints() {
	//This define a constant during the program
	//It is not necesary set the var value
	const winner string = "I am a winner"

}

//Declaring multiples variables
func multipleVariables() {
	var (
		nickName  string
		edad      int
		birthdate string
	)
	nickName = "fibannaci"
	edad = 20
	birthdate = "21/06/2016"
	fmt.Println("The nickname is ", nickName)
	fmt.Println("The age is ", edad)
	fmt.Println("The birthdate is ", birthdate)
}
