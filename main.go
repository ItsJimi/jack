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
	Port int    `json:"port"`
	Path string `json:"path"`
}

func main() {
	c := &config{8080, "."}

	configFile, err := ioutil.ReadFile(".serve.json")
	if err == nil {
		err = json.Unmarshal(configFile, &c)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}

	port := flag.Int("port", c.Port, "Port")
	path := flag.String("path", c.Path, "Directory")

	flag.Parse()

	if *port <= 0 || *port >= 65535 {
		log.Fatalf("Invalid port")
	}

	fmt.Printf("Start on %d\n", *port)
	http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(*path)))
}
