# Go Webserver Live Coding 

## Schritt 0: Initiales Setup

### a) Go Modul initialisieren
```
go mod init java-to-go
```

### b) Datei `main.go` erstellen
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Congressman!")
}
```


## Schritt 1

### `main.go`
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Congressman!")
	})
	http.ListenAndServe(":8090", nil)
}
```

## Schritt 2

### `main.go`
```go
package main

import (
	"fmt"
	"net/http"
)

func handleCongressman(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Congressman!")
}

func main() {
	http.HandleFunc("/", handleCongressman)
	http.ListenAndServe(":8090", nil)
}
```

## Schritt 2

### `congressman.go`
```go
package main

type Congressman struct {
	Id   string
	Name string
}
```

### `main.go`
```go
package main

import (
	"encoding/json"
	"net/http"
)

func handleCongressman(w http.ResponseWriter, r *http.Request) {
	c := Congressman{Id: "pr", Name: "Peter Russo"}
	json.NewEncoder(w).Encode(c)
}

func main() {
	http.HandleFunc("/", handleCongressman)
	http.ListenAndServe(":8090", nil)
}
```

## Schritt 3

### `main.go`
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var congress []Congressman

func handleCongressman(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	for i, c := range congress {
		if c.Id == id {
			log.Printf("Found at %v", i)
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	fmt.Fprintf(w, "- not found -")
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	congress = make([]Congressman, 2)
	congress[0] = Congressman{Id: "pr", Name: "Peter Russo"}
	congress[1] = Congressman{Id: "fu", Name: "Frank Underwood"}

	http.HandleFunc("/", handleCongressman)
	http.ListenAndServe(":8090", nil)
}
```