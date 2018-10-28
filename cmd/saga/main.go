package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tombell/saga"
)

const helpText = `usage: saga [options] session file...

Special options:
  --help     show this message, then exit
  --version  show the version number, then exit
`

var (
	version = flag.Bool("version", false, "")
)

func usage() {
	fmt.Fprintf(os.Stderr, helpText)
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *version {
		fmt.Fprintf(os.Stdout, "saga %s (%s)\n", Version, Commit)
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
	}

	logger := log.New(os.Stderr, "[saga] ", log.LstdFlags)

	cfg := saga.Config{
		Filepath: args[0],
		Logger:   logger,
	}

	if err := saga.Run(cfg); err != nil {
		panic(err)
	}
}
