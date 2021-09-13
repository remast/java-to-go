package main

import (
	"fmt"
	"time"
)

func HelloCongressman(name string) {
	fmt.Println("Hello Congressman", name)
}

func main() {
	HelloCongressman("Russo")

	time.Sleep(5 * time.Millisecond)
}
