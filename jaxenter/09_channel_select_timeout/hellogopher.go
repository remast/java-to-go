package main

import (
	"fmt"
	"time"
)

func Hello(names chan string) {
	select {
	case name := <-names:
		fmt.Println("Hello " + name + "!")
	case <-time.After(time.Second):
		fmt.Println("No one here ...!")
	}
}

func main() {
	names := make(chan string)
	go Hello(names)
	time.Sleep(3 * time.Second)
}
