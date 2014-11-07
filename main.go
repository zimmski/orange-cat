package main

import (
	"github.com/codegangsta/cli"
	"os"
)

const Version = "0.2.1-dev"

func main() {
	app := cli.NewApp()
	app.Name = "orange"
	app.Version = Version
	app.Usage = `orange-cat is a Markdown previewer written in Go.
   Its main goal is to be used with any editor you love.
   For information, please visit https://github.com/noraesae/orange-cat`
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "basic, b",
			Usage: "Use Markdown Basic(Markdown Common by default).",
		},
		cli.IntFlag{
			Name:  "port, p",
			Value: 6060,
			Usage: "Port to listen.",
		},
	}
	app.Action = func(c *cli.Context) {
		args := c.Args()

		orange := NewOrange(c.Int("port"))

		if c.Bool("basic") {
			orange.UseBasic()
		}

		orange.Run(args...)
	}

	// codegangsta/cli help template
	cli.AppHelpTemplate = `orange-cat
   {{.Usage}}

USAGE:
   {{.Name}} [global options] [command] file

COMMANDS:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
   {{end}}{{if .Flags}}
GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
`

	app.Run(os.Args)
}
