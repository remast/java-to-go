package main

import "fmt"

type Congressman struct {
	Name string
}

func main() {
	c := Congressman{Name: "Peter Russo"}
	fmt.Println("Hello " + c.Name + "!")
}
