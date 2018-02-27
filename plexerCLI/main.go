package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

const (
	version = "0.0.1"
	verbose = false
)

var (
	plexAddr  string
	plexToken string
)

func main() {

	app := cli.NewApp()
	app.Name = "plexer"
	app.Version = version
	app.Usage = "Perform Plex stuff"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "plexAddr",
			Value:       "",
			Usage:       "Plex IP Address",
			Destination: &plexAddr,
			EnvVar:      "PLEX_ADDR",
		},
		cli.StringFlag{
			Name:        "token",
			Value:       "",
			Usage:       "Plex Token",
			Destination: &plexToken,
			EnvVar:      "PLEX_TOKEN",
		},
	}

	app.Commands = []cli.Command{
		// listCommand(),
		// searchCommand(),
		// restoreCommand(),
		playlistCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

func logXML(data interface{}) {
	output, err := xml.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%s\n", output)
}
