package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/thoas/observr/application"
	"github.com/thoas/observr/config"
	"github.com/thoas/observr/events"
	"github.com/thoas/observr/worker"
)

func main() {
	app := cli.NewApp()
	app.Name = "observr"
	flags := []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			EnvVar: "OBSERVR_CONF",
			Usage:  "Configuration file",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "Start application",
			Flags:   flags,
			Action: func(c *cli.Context) {
				_, err := config.Load(c.String("config"))

				if err != nil {
					panic(err)
				}
			},
		},
		{
			Name:    "worker",
			Aliases: []string{"w"},
			Usage:   "Start workers",
			Action: func(c *cli.Context) {
				worker, err := worker.Load(c.String("config"))

				if err != nil {
					panic(err)
				}

				worker.Run()
			},
		},
		{
			Name: "producer",
			Action: func(c *cli.Context) {
				ctx, err := application.Load(c.String("config"))

				if err != nil {
					panic(err)
				}

				events := events.FromContext(ctx)

				events.Producer.Publish("test", []byte(`{"foo": "bar"}`))
			},
		},
	}

	app.Run(os.Args)
}
