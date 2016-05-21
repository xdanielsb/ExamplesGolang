package main

import "fmt"

func main() {
	d1 := Dwarf{name: "Gimli"}
	d2 := Dwarf{name: "Tyrion"}

	e1 := Elf{name: "Legolas"}
	e2 := Elf{name: "Daniel"}

	fighting(d1, d2, e1, e2)
}

type Personaje interface {
	fight()
}

type Dwarf struct {
	name string
}

type Elf struct {
	name string
}

func (d Dwarf) fight() {
	fmt.Printf("I am %s fighting with my maze \n", d.name)
}

func (e Elf) fight() {
	fmt.Printf("I am %s fighting with my arc \n", e.name)
}

func fighting(warriors ...Personaje) {
	for _, warrior := range warriors {
		warrior.fight()
	}
}
