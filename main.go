package main

import (
	"bufio"
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

func serve(ctx *cli.Context) error {
	var c config

	configFile, err := ioutil.ReadFile(ctx.GlobalString("config"))
	if err == nil {
		err = json.Unmarshal(configFile, &c)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}

	if c.Path != "" {
		ctx.Set("path", c.Path)
	}
	if c.Addr != "" {
		ctx.Set("addr", c.Addr)
	}
	if c.Port != 0 {
		ctx.Set("port", strconv.Itoa(c.Port))
	}

	if ctx.Int("port") <= 1024 || ctx.Int("port") >= 65535 {
		log.Fatalf("Invalid port")
	}

	srv := &http.Server{
		Addr:    ctx.String("addr") + ":" + ctx.String("port"),
		Handler: http.FileServer(http.Dir(ctx.String("path"))),
	}
	fmt.Printf("Start on %s:%d\n", ctx.String("addr"), ctx.Int("port"))
	http2.ConfigureServer(srv, &http2.Server{})
	log.Fatal(srv.ListenAndServe())

	return nil
}

func connect(ctx *cli.Context) error {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for {
		fmt.Print("> ")
		scanner.Scan()
		fmt.Println(scanner.Text())
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "Jack"
	app.Usage = "🏗  Amazing web development tool"
	app.Version = "0.1.1"

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
			Action:  connect,
		},
	}

	app.Run(os.Args)
}
