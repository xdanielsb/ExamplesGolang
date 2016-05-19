// Control Structures

package main

import "fmt"

func main() {
	forStatement()
	ifStatement()
	switchStatement()
}

func forStatement() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}
	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}
}

func ifStatement() {
	var number1 = 541
	var number2 = 84

	if number1 == number2 {
		fmt.Println("the numbers are the same")
	} else if number1 > number2 {
		fmt.Println("the number  number1  es mayor que el numero2")
	} else {
		fmt.Println("the number  number2  es mayor que el numero1")
	}
}

func switchStatement() {
	var number int64 = 5
	fmt.Println("Please user write a number between [0 5]")
	fmt.Scanf("%i", &number)
	fmt.Println("The number is: ", number)
	switch number {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("The number is not in the range")
	}
}
