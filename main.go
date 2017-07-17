package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type config struct {
	Port int64  `yaml:"port"`
	Path string `yaml:"path"`
}

func main() {
	c := config{}

	yamlFile, err := ioutil.ReadFile(".serve.yml")
	if err != nil {
		log.Print(err)
		return
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Print(yamlFile)
	// port := flag.Int("port", 8080, "Port")
	// path := flag.String("path", ".", "Directory")

	// flag.Parse()

	// fmt.Printf("Start on %d\n", *port)
	// http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(*path)))
}
