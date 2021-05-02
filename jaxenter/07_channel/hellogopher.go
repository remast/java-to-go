package main

import "fmt"

func Hello(names chan string) {
	name := <-names
	fmt.Println("Hello " + name + "!")
}

func main() {
	names := make(chan string)
	go Hello(names)
	names <- "Jim"
}
