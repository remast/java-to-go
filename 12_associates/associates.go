package main

import (
	"fmt"
	"math/rand"
	"time"
)

func associate(name string, jobs <-chan int, results chan<- string) {
	for job := range jobs {
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		fmt.Println(name, "finished job", job)
		results <- fmt.Sprintf("Draft Law §%v", job)
	}
}

func main() {
	jobs := make(chan int, 10)       // Channel für Jobs
	results := make(chan string, 10) // Channel für Ergebnisse

	go associate("John", jobs, results)  // Worker 1
	go associate("Maria", jobs, results) // Worker 2

	for j := 1; j <= 5; j++ { // 5 Jobs einstellen
		jobs <- j
	}
	close(jobs) // Channel schließen

	for a := 1; a <= 5; a++ { // Ergebnisse abholen
		fmt.Printf("Received '%v'\n", <-results) // HL
	}
}
