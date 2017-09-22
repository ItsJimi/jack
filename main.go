package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/net/http2"
)

type config struct {
	Addr string `json:"addr"`
	Port int    `json:"port"`
	Path string `json:"path"`
}

func main() {
	c := &config{"", 8080, "."}

	configFile, err := ioutil.ReadFile(".serve.json")
	if err == nil {
		err = json.Unmarshal(configFile, &c)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}

	addr := flag.String("addr", c.Addr, "Address")
	port := flag.Int("port", c.Port, "Port")
	path := flag.String("path", c.Path, "Directory")

	flag.Parse()

	if *port <= 0 || *port >= 65535 {
		log.Fatalf("Invalid port")
	}

	srv := &http.Server{
		Addr:    *addr + ":" + strconv.Itoa(*port),
		Handler: http.FileServer(http.Dir(*path)),
	}
	fmt.Printf("Start on %s:%d\n", *addr, *port)
	http2.ConfigureServer(srv, &http2.Server{})
	log.Fatal(srv.ListenAndServe())
}
