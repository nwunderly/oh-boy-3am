package main

import (
	"context"
	"github.com/akamensky/argparse"
	"github.com/nwunderly/oh-boy-3am/bot"
	"github.com/nwunderly/oh-boy-3am/db"
	"os"
)

type Args struct {
	Dev  bool
	Debug bool
}

func parseArgs() Args {
	parser := argparse.NewParser("Oh boy, 3AM!", "The most useful discord bot.")
	dev := parser.Flag("", "dev", &argparse.Options{
		Required: false,
		Help:     "Whether to use production token or run as dev bot. Production also uses a local database connection.",
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

	return Args{Dev: *dev, Debug: *debug}
}

func main() {
	args := parseArgs()

	var token string
	var postgresURL string

	if args.Dev {
		token = tokenDev
		postgresURL = postgresExternalURL
	} else {
		token = tokenProd
		postgresURL = postgresInternalURL
	}

	b := bot.New(":=", token, args.Debug)

	err := db.Database.Connect(postgresURL)
	if err != nil {
		panic(err)
	}
	defer db.Database.Conn.Close(context.Background())

	b.Run()
}
