package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{ message: 'Hey !' }")
	fmt.Print(r)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{ test: true }")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/test", testHandler)
	http.ListenAndServe(":8080", nil)
}
