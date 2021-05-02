package main

import (
	"fmt"
	"sync"
)

func HelloGopher() {
	fmt.Println("Hello Gopher!")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		HelloGopher()
	}()
	wg.Wait()
}
