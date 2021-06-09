package main

import (
	"os"
)

func main() {
	os.Remove("10_webserver/main.go")
	os.Remove("10_webserver/congressman.go")
}
