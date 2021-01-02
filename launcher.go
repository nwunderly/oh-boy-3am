package main

import (
	"github.com/akamensky/argparse"
	"github.com/nwunderly/oh-boy-3am/bot"
	"os"
)

type Args struct {
	Prod  bool
	Debug bool
}

func parseArgs() Args {
	parser := argparse.NewParser("Oh boy, 3AM!", "The most useful discord bot.")
	prod := parser.Flag("", "prod", &argparse.Options{
		Required: false,
		Help:     "Whether to use production token or run as dev bot. Defaults to",
		Default:  false,
	})
	debug := parser.Flag("", "debug", &argparse.Options{
		Required: false,
		Help:     "Whether to log in debug mode. Log level defaults to Info otherwise.",
		Default:  false,
	})

	err := parser.Parse(os.Args)
	if err != nil {
		panic(err)
	}

	return Args{Prod: *prod, Debug: *debug}
}

func main() {
	args := parseArgs()

	var token string
	if args.Prod {
		token = tokenProd
	} else {
		token = tokenDev
	}

	b := bot.New(":=", token, args.Debug)

	b.Run()
}
