package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tombell/saga"
)

const helpText = `usage: saga [options]

Saga options:
  --session-dir   path to the sessions directory to watch for the next session file
  --session-file  path to the session file
  --listen        host/port to listen on

Special options:
  --help          show this message, then exit
  --version       show the version number, then exit
`

var (
	sessionFile = flag.String("session-file", "", "")
	sessionDir  = flag.String("session-dir", "", "")
	listen      = flag.String("listen", ":8080", "")
	version     = flag.Bool("version", false, "")
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

	if *sessionFile != "" && *sessionDir != "" {
		fmt.Fprintln(os.Stderr, "cannot use both --session-dir and --session-file\n")
		flag.Usage()
	}

	if *sessionFile == "" && *sessionDir == "" {
		fmt.Fprintln(os.Stderr, "must use either --session-dir or --session-file\n")
		flag.Usage()
	}

	logger := log.New(os.Stderr, "[saga] ", log.LstdFlags)

	cfg := saga.Config{
		Logger:      logger,
		Listen:      *listen,
		SessionDir:  *sessionDir,
		SessionFile: *sessionFile,
	}

	if err := saga.Run(cfg); err != nil {
		panic(err)
	}
}
