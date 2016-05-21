//goroutines:: is a function that is capable of run concurrently with other funcitons
package main

import "fmt"

func test(n int) {
	for i := 1; i <= 1000; i++ {
		fmt.Printf("%d : %d \n", n, i)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go test(i)
	}
	var input string
	fmt.Scanln(&input)
}
