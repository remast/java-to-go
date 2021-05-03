package main

import "fmt"

type Congressman struct {
	Name string
}

func (c Congressman) swearOathOfOffice() {
	fmt.Printf("I, %v, swear to serve the USA.", c.Name)
}

type President struct {
	// Embedded
	Congressman

	NuclearWeaponCode string
}

func main() {
	p := President{NuclearWeaponCode: "123"}
	p.Name = "Frank Underwood"
	p.swearOathOfOffice()
}
