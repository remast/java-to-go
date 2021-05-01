package main

import (
	"fmt"
	"sync"
)

func HelloCongressman(name string) {
	fmt.Println("Hello Congressman", name)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		HelloCongressman("Russo")
	}()
	wg.Wait()
}
