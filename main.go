package main

import (
	"github.com/codegangsta/cli"
	"github.com/thoas/observr/application"
	"os"
)

func main() {
	application, err := application.New()

	if err != nil {
		panic(err)
	}

	app := cli.NewApp()
	app.Name = "observr"

	app.Commands = []cli.Command{
		{
			Name:    "worker",
			Aliases: []string{"w"},
			Usage:   "Start workers",
			Action: func(c *cli.Context) {
				application.Work()
			},
		},
		{
			Name: "producer",
			Action: func(c *cli.Context) {
				application.EventStore.Producer.Publish("test", []byte(`{"foo": "bar"}`))
			},
		},
	}

	app.Run(os.Args)
}
