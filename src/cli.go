package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

type CLI struct {
	app *cli.App
}

func (e *ego) DoCLI() {
	app := &cli.App{
		Name: "Ego",
		Action: func(c *cli.Context) error {
			e.DoRender()
			if c.IsSet("p"){
				q := NewQiniu()
				q.Upload()
			}

			if c.String("s") != "" {
				e.DoServer(c.String("s"))
			}
			return nil
		},
	}
	e.CLI = new(CLI)
	e.CLI.app = app
	e.CLI.setFlag()

	err := e.CLI.app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *CLI) setFlag() {
	c.app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "s",
			Usage:    "启动 HTTP 服务器",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "p",
			Usage:    "发布",
			Required: false,
		},
	}
}
