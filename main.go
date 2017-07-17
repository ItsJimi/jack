package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type config struct {
	Port int64  `json:"port"`
	Path string `json:"path"`
}

func main() {
	c := config{}

	configFile, err := ioutil.ReadFile(".serve.json")
	if err != nil {
		log.Print(err)
		return
	}
	err = json.Unmarshal(configFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Print(c.Path)

	port := flag.Int("port", 8080, "Port")
	path := flag.String("path", ".", "Directory")

	flag.Parse()

	fmt.Printf("Start on %d\n", *port)
	http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(*path)))
}
