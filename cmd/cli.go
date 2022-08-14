package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"projecterror.dev/cli/internal/resource"
)

	
var (
	boilerplateName string
	boilerplateLanguage string
)


func main() {
	app := &cli.App{
		Name:  "pecli",
		Usage: "Create npwd templates, or boilerplates",
		Commands: []*cli.Command{
			{
				Name:    "boilerplate",
				Aliases: []string{"b"},
				Usage:   "Configure a new boilerplate",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "name",
						Usage:       "Name of the resource",
						Destination: &boilerplateName,
						Required: true,
					},
					&cli.StringFlag{
						Name:        "langauge",
						Aliases:     []string{"l", "lang"},
						Usage:       "Language to use",
						Destination: &boilerplateLanguage,
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					if c.String("langauge") == "lua" {
						resource.CreateLuaResource(boilerplateName)
					}

					if c.String("langauge") == "javascript" {
						resource.CreateJSResource(boilerplateName)
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
