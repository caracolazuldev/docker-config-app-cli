package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type Container struct {
	Files iFileSys
	Term  iUser
}

func initServices() *Container {
	return &Container{
		Files: &FileSys{},
		Term:  &Terminal{},
	}
}

var services *Container

func main() {

	services = initServices()

	app := &cli.App{
		Name: "add-config",
		// Flags: []cli.BoolFlag{
		// 	Name: "-q"
		// 	Usage: "quiet mode",
		// },
		Action: addConfig,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
