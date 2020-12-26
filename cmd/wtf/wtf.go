package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "WSDL To Formly (WTF)",
		Usage: "Convert a webservice (WSDL) specification to angular formly input schema",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Peter A. Ramaldes",
				Email: "peter.ramaldes@gmail.com",
			},
		},
		Version: "0.0.1",

		Action: func(c *cli.Context) error {
			fmt.Println("WSDL:", c.Args().Get(0))
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
