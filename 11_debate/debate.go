package main

import (
	"fmt"
	"time"
)

func randomAnswer() string {
	return "Nooo!!"
}

func speaker(name string, debate chan int) {
	for {
		microphone := <-debate // Auf Mikro warten

		fmt.Printf("Turn %v: %v says '%v'\n", microphone, name, randomAnswer())
		time.Sleep(400 * time.Millisecond)

		microphone++
		debate <- microphone // Mikro zurückgeben
	}
}

func main() {
	debate := make(chan int)

	go speaker("Jackie", debate) // Kandidat 1
	go speaker("Frank", debate)  // Kandidat 2

	microphone := 1
	debate <- microphone        // Mikro geben und Diskussion starten
	time.Sleep(2 * time.Second) // Dauer der Diskussion*
	<-debate                    // Mikro nehmen und Diskussion beenden
}
