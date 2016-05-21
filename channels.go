//Channels is a way to communicate two go routines #AWESOME
//the message is send and is receive by the operator <-

package main

import (
	"fmt"
	"time"
)

/*

This means that c can only be sent
func pinger(c chan<- string) {

This means that c can only receive
func pinger(c <-chan string) {

This means that the channel can receive and sent because does not have restrictions
func pinger(c chan string) {


*/

//this a sender : sender number 1
func pinger(c chan string) {
	for i := 0; ; i++ {
		//this part send the message
		c <- "ping"
	}
}

//this a sender : sender number 1
func ponger(c chan string) {
	for i := 0; ; i++ {
		//this part send the message
		c <- "pong"
	}
}

//this is the listener
func printer(c chan string) {
	for {
		//this part <- receives the message
		msg := <-c
		fmt.Printf("%s \n", msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {

	var c chan string = make(chan string)
	go pinger(c)  //this  sentence called the function that assign the message //this is a the sender 1
	go ponger(c)  //this  sentence called the function that assign the message //this is a the sender 2
	go printer(c) //this sentence called the function that receive the message //this is the listener

	var input string
	fmt.Scanln(&input)

}
