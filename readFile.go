package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs := open("in.txt")
	str := string(bs)
	fmt.Printf("%s", str)
}

func open(name string) (bs []byte) {
	bs, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Printf("There was an error during the execution\n")
		return
	}
	return
}
