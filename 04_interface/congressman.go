package main

import "fmt"

type Congressman struct {
	Name string
}

func (c Congressman) greet() {
	fmt.Println("Hello", c.Name)
}

type Enemy struct{}

func (e Enemy) greet() {
	fmt.Println("Go to hell!")
}

type Greeter interface {
	greet()
}

func passBy(c1 Greeter, c2 Greeter) {
	c1.greet()
	c2.greet()
}

func main() {
	c := Congressman{Name: "Frank U."}
	e := Enemy{}
	passBy(c, e)
}
