package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"projecterror.dev/cli/pkg/commands/boilerplate"
)

	

func main() {
	app := &cli.App{
		Name:  "pecli",
		Usage: "Create npwd templates, or boilerplates",
		Commands: []*cli.Command{
			boilerplate.RootCmd(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
