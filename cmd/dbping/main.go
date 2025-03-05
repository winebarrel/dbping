package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/winebarrel/dbping"
)

var version string

func parseArgs() *dbping.Config {
	var CLI struct {
		dbping.Config
		Version kong.VersionFlag
	}

	parser := kong.Must(&CLI, kong.Vars{"version": version})
	parser.Model.HelpFlag.Help = "Show help."
	_, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)

	return &CLI.Config
}

func main() {
	config := parseArgs()
	dbping.Ping(config)
}
