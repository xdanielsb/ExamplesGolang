//Maps = key value pairs = asociative arrays = hash table = dictionaries

package main

import "fmt"

func main() {
	creatingMap()
	deleteItems()
	creatingMap2()
	mapsMaps()
}

//First Form
func creatingMap() {
	var users map[string]int
	//before uses the maps is necesary initialize them
	users = make(map[string]int)

	//Now we can use the map
	users["Ana"] = 8
	fmt.Println(users["Ana"])
}

//Second Form
func creatingMap2() {
	users := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		"E": 4,
	}
	fmt.Println(users)
}

//Creating Maps  of maps
func mapsMaps() {
	users := map[string]map[string]int{
		"A": map[string]int{"AX": 1, "AXO": 2},
		"B": map[string]int{"BX": 1, "BXO": 2},
		"C": map[string]int{"CX": 1, "CXO": 2},
		"D": map[string]int{"DX": 1, "DXO": 2},
		"E": map[string]int{"EAX": 1, "EXO": 2},
	}
	fmt.Println(users)
}

func deleteItems() {
	var calificaciones map[string]int
	//initialize the array
	calificaciones = make(map[string]int)

	//adding the values
	calificaciones["Daniel 0"] = 50
	calificaciones["Daniel 1"] = 40
	calificaciones["Daniel 2"] = 30
	calificaciones["Daniel 3"] = 30
	calificaciones["Daniel 4"] = 20

	delete(calificaciones, "Daniel 0")

	//if exist is true means that the key are in the dictionary
	grade, exist := calificaciones["Daniel 0"]
	fmt.Println(grade, exist)

	//Much beatiful form to do that
	//first we try to get the value from the map then if it successfull we run the code inside the block
	if grade, exist := calificaciones["Daniel 1"]; exist {
		fmt.Println(grade, exist)
	}

}
