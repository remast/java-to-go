package main

import "fmt"

type Person struct {
	Name string
}

func nameFritz(p *Person) {
	p.Name = "Fritz"
}

func main() {
	p := Person{Name: "Max"}
	nameFritzPointer(&p)
	fmt.Println("Hello " + p.Name)
}
