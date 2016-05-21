package main

import "fmt"

type Car struct {
	placa string
}

type Person struct {
	name string
	id   int
	car  Car //This type of relation is called has-a  : Composition  : Person has a car
}

type Boss struct {
	Person        //This type of relations is called is a : inheritante: Boss is a Person
	numeroOficina int
}

func (c Car) getPlaca() (placa string) {
	placa = c.placa
	return
}

func (p Person) getName() (name string) {
	name = p.name
	return
}

func (p Person) getCar() (car Car) {
	car = p.car
	return
}

func main() {
	car1 := Car{placa: "PDS85"}
	person1 := Person{name: "Daniel", id: 25, car: car1}
	boss1 := Boss{Person: person1, numeroOficina: 5}

	//Now, we are going to show how it works
	//Composition

	fmt.Printf("%s \n", person1.car.getPlaca())

	//Inheritance
	fmt.Printf("%s \n", boss1.getName())
}
