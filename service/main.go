package main

import (
	"log"
	"new-service/application"
	"new-service/config"
	"new-service/route"
	"os"

	"github.com/urfave/cli"
)

func main() {
	cfg, err := config.Setup()
	if err != nil {
		log.Fatal("Cannot load config ", err.Error())
	}

	if Cli(cfg).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func Cli(cfg *config.Config) *cli.App {
	clientApp := cli.NewApp()
	clientApp.Name = cfg.Application.ServiceName
	clientApp.Action = runCommand(cfg)
	clientApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "mode",
			Required: false,
			Value:    "grpc",
		},
	}
	return clientApp
}

func runCommand(cfg *config.Config) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		app, err := application.Setup(cfg, c)
		if err != nil {
			return err
		}
		route.SetupGrpcRouter(cfg, app)
		return app.Run(cfg)
	}
}
