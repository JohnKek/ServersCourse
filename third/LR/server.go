package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello/", hello)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := strings.Split(r.URL.Path, "/")[2]
	w.Write([]byte(fmt.Sprintf("Hello, %s", name)))
}
