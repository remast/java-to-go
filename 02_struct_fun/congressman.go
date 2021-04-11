package main

import "fmt"

type Congressman struct {
	Name string
}

func (c Congressman) swearOathOfOffice() {
	fmt.Printf("I, %v, swear to serve the USA.", c.Name)
}

func main() {
	c := Congressman{Name: "Peter Russo"}
	c.swearOathOfOffice()
}
