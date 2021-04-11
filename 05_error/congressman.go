package main

import (
	"errors"
	"fmt"
)

type Congressman struct {
	Name           string
	AccountBalance float64
}

func (c Congressman) bribe(amount float64) error {
	if c.Name != "Peter Russo" {
		return errors.New("Not corrupt!")
	}
	c.AccountBalance += amount
	return nil
}

func main() {
	c := Congressman{Name: "Jackie Sharp", AccountBalance: -10.0}
	err := c.bribe(5000.0)

	if err != nil {
		fmt.Printf("%v is not bribable.\n", c.Name)
	}

	fmt.Printf("%v is worth %v", c.Name, c.AccountBalance)
}
