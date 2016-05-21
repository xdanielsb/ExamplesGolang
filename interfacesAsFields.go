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
	speak()
}

type Dwarf struct {
	name string
}

type Elf struct {
	name string
}

func (d Dwarf) fight() {
	fmt.Printf("I %s am fighting with my maze", d.name)
}

func (e Elf) fight() {
	fmt.Printf("I %s am fighting with my arc", d.name)
}

func fighting(warriors ...Personaje) {
	for _, warrior := range warriors {
		warrior.fight()
	}
}
