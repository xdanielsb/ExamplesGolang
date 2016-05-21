package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		go delay(i)
	}
	var in string
	fmt.Scanln(&in)
}

func delay(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d : %d \n", n, i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
