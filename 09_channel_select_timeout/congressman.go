package main

import (
	"fmt"
	"time"
)

func Congressman(money chan int) {
	// Nachricht empfangen
	select {
	case amount := <-money:
		fmt.Println("Received", amount, "$!")
	case <-time.After(1 * time.Second):
		fmt.Println("Got nothing ...!")
	}
}

func main() {
	money := make(chan int)
	go Congressman(money)

	time.Sleep(2 * time.Second)
}
