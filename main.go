package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/thoas/observr/application"
	"github.com/thoas/observr/broker"
	"github.com/thoas/observr/web"
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
				ctx, err := application.Load(c.String("config"))

				if err != nil {
					panic(err)
				}

				web.Run(ctx)
			},
		},
		{
			Name:    "worker",
			Aliases: []string{"w"},
			Usage:   "Start workers",
			Flags:   flags,
			Action: func(c *cli.Context) {
				ctx, err := application.Load(c.String("config"))

				if err != nil {
					panic(err)
				}

				worker.Run(ctx)
			},
		},
		{
			Name:    "producer",
			Aliases: []string{"p"},
			Flags:   flags,
			Action: func(c *cli.Context) {
				ctx, err := application.Load(c.String("config"))

				if err != nil {
					panic(err)
				}

				b := broker.FromContext(ctx)
				b.Publish(ctx, &broker.UserCreatedEvent{Username: "thoas"})
			},
		},
	}

	app.Run(os.Args)
}
