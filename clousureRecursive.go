//Create Functions inside functions

package main

import "fmt"

func main() {
	increment()
	suma()
	suman()
	fmt.Println(factorial(1400))
}

//Function increment inside a function
func increment() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

func suma() {
	var (
		num1 float64
		num2 float64
	)
	fsuma := func(x float64, y float64) (resultado float64) {
		resultado = y + x
		return
	}
	num1 = 652
	num2 = 58
	fmt.Println(fsuma(num1, num2))

}

func suman() {
	fsuma := func(values ...float64) (resultado float64) {
		resultado = 0
		for _, value := range values {
			resultado += value
		}
		return
	}
	fmt.Println(fsuma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

//Now we are going to do some recursive functions

func factorial(n float64) (result float64) {
	if n == 0 {
		result = 1
	} else {
		result = n * factorial(n-1)
	}
	return
}
