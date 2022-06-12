package main

import (
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"projecterror.dev/cli/internal/resource"
)

func main() {
	var boilerplateName string
	var boilerplateLanguage string

	app := &cli.App{
		Name:  "pe",
		Usage: "Create npwd templates, or boilerplates",
		Commands: []*cli.Command{
			{
				Name:    "boilerplate",
				Aliases: []string{"b"},
				Usage:   "Configure a new boilerplate",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "name",
						Aliases:     []string{"n"},
						Usage:       "Name of the resource",
						Destination: &boilerplateName,
					},
					&cli.StringFlag{
						Name:        "langauge",
						Aliases:     []string{"l", "lang"},
						Usage:       "Language to use",
						Destination: &boilerplateLanguage,
					},
				},
				Action: func(c *cli.Context) error {
					cRed := color.New(color.FgGreen)
					cRed.Println("Configure a new boilerplate")

					if c.String("langauge") == "lua" {
						resource.CreateLuaResource(boilerplateName)
					}

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
