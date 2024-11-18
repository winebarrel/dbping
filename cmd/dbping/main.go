package main

import (
	"log"
	"os"

	"github.com/alecthomas/kong"
	"github.com/winebarrel/dbping"
)

var version string

func init() {
	log.SetFlags(0)
}

func parseArgs() *dbping.DBConfig {
	var CLI struct {
		dbping.DBConfig
		Version kong.VersionFlag
	}

	parser := kong.Must(&CLI, kong.Vars{"version": version})
	parser.Model.HelpFlag.Help = "Show help."
	_, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)

	return &CLI.DBConfig
}

func main() {
	config := parseArgs()
	dbping.Ping(config)
}
