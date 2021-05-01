package main

import (
	"fmt"
	"time"
)

func Congressman(money chan int) {
	// Nachricht empfangen
	amount := <-money
	fmt.Println("Received", amount, "$!")
}

func main() {
	money := make(chan int)
	go Congressman(money)

	// Nachricht senden
	money <- 100

	time.Sleep(5 * time.Millisecond)
}
