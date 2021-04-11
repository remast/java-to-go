package main

import "fmt"

type Congressman struct {
	Name           string
	AccountBalance float64
}

func (c *Congressman) bribe(amount float64) {
	c.AccountBalance += amount
}

func main() {
	c := &Congressman{Name: "Peter Russo", AccountBalance: -10.0}
	c.bribe(5010.0)

	fmt.Printf("%v is worth %v", c.Name, c.AccountBalance)
}
