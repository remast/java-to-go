package main

import "fmt"

type Person struct {
	Name string
}

type Stepmother struct{}

func (p Stepmother) greet() {
	fmt.Println("Go to hell!")
}

func passBy(p1 NicePerson, p2 NicePerson) {
	p1.greet()
	p2.greet()
}

type NicePerson interface {
	greet()
}

func (p Person) greet() {
	fmt.Println("Hello " + p.Name + "!")
}

func main() {
	p := Person{Name: "j"}
	s := Stepmother{}
	passBy(p, s)
}
