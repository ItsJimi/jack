package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 8080, "Port")
	path := flag.String("path", ".", "Directory")

	flag.Parse()

	fmt.Printf("Start on %d\n", *port)
	http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(*path)))
}
