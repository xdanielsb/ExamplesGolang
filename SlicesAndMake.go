//Tests with Slices

package main

import "fmt"

//Declaring a global variable

func main() {
	slices()
	makeArrays()
}

func slices() {
	array := [10]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array)

	//printing a part of an array
	fmt.Println("Subarray ", array[1:4])

	slice1 := []float64{1, 2, 3, 4, 5}
	fmt.Println("Slice 1: ", slice1)

	//appending to and slice
	slice2 := append(slice1, 43, 234, 123, 12)
	fmt.Println("Slice 2: ", slice2)

	//using make sentence
	//10 represent the number of elements of the array
	slice3 := make([]float64, 5, 10)
	fmt.Println("Slice 3: ", slice3)

	//copyng and slice in another slice
	//copy the slice 2 in the slice 1
	copy(slice1, slice2)
	fmt.Println("Slice 1: ", slice1)

}

//Make save the necesary space for an array
func makeArrays() {
	fmt.Println("Making arrays")
	array := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array)
	slice := make([]float64, 4)
	fmt.Println(slice)
	//We are going to copy the array in to the slice

	copy(slice, array) //Just copy 4 element of the array because the size is 4
	fmt.Println(slice)

}
