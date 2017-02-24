package main

import "testing"

func TestSayHelloGopher(t *testing.T) {
	if HelloGopher() != "Hello Gopher!" {
		t.Error("Unerwartetes Ergebnis.")
	}
}
