//Structs

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	structExample()
}

func structExample() {
	//Create an instance form 1
	var person1 *Person
	//Initialice the instance of cereate the instance
	person1 = new(Person)

	//Create an instance form 2
	person2 := Person{name: "Daniel", age: 19}

	//Create an instance form 3
	//of this way it is necesary add the elements in the order who appears uin the struct
	person3 := Person{"Daniel", 19}

	fmt.Println(person1.name)
	fmt.Println(person2.name)
	fmt.Println(person3.name)
}
