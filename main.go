package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/urfave/cli"
	"golang.org/x/net/http2"
)

type config struct {
	Addr string `json:"addr"`
	Port int    `json:"port"`
	Path string `json:"path"`
}

func serve(cliContext *cli.Context) error {
	var c config

	configFile, err := ioutil.ReadFile(cliContext.GlobalString("config"))
	if err == nil {
		err = json.Unmarshal(configFile, &c)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}

	path := cliContext.String("path")
	addr := cliContext.String("addr")
	port := cliContext.Int("port")

	if port <= 0 || port >= 65535 {
		log.Fatalf("Invalid port")
	}

	srv := &http.Server{
		Addr:    addr + ":" + strconv.Itoa(port),
		Handler: http.FileServer(http.Dir(path)),
	}
	fmt.Printf("Start on %s:%d\n", addr, port)
	http2.ConfigureServer(srv, &http2.Server{})
	log.Fatal(srv.ListenAndServe())

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "jack"
	app.Usage = "amazing tool for web development"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: ".jack.json",
			Usage: "path of config file",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "serve",
			Aliases: []string{"s"},
			Usage:   "serve static files",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "path",
					Value: ".",
					Usage: "path of static files",
				},
				cli.StringFlag{
					Name:  "addr",
					Value: "0.0.0.0",
					Usage: "address of server",
				},
				cli.StringFlag{
					Name:  "port",
					Value: "8080",
					Usage: "port of server",
				},
			},
			Action: serve,
		},
		{
			Name:    "connect",
			Aliases: []string{"c"},
			Usage:   "connect to websocket server",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	app.Run(os.Args)
}
