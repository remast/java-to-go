package main

import "fmt"

func Hello(names chan string) {
	select {
	case name := <-names:
		fmt.Println("Hello " + name + "!")
	}
}

func main() {
	names := make(chan string)
	go Hello(names)
	names <- "Jim"

	names <- "Joe"
	fmt.Println("Done")
}
