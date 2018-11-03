package saga

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tombell/saga/decks"
	"github.com/tombell/saga/monitor"
	"github.com/tombell/saga/server"
	"github.com/tombell/saga/watcher"
)

// Config ...
type Config struct {
	Logger      *log.Logger
	Listen      string
	SessionDir  string
	SessionFile string
}

// Run begins the process of listening for changes to the given Serato session
// file. It keeps a realtime model of the decks, and the tracks they're
// playing, and have played.
func Run(cfg Config) error {
	if cfg.SessionDir != "" {
		cfg.Logger.Printf("waiting for new session in %s...\n", cfg.SessionDir)

		file, err := watcher.WaitForSession(cfg.SessionDir)
		if err != nil {
			return err
		}

		cfg.SessionFile = file
	}

	decks := decks.NewDecks(cfg.Logger)

	monitor, err := monitor.New(monitor.Config{
		Logger:   cfg.Logger,
		Decks:    decks,
		Filepath: cfg.SessionFile,
	})
	if err != nil {
		return err
	}
	defer monitor.Close()

	monitorErrCh := make(chan error, 1)
	go monitor.Run(monitorErrCh)

	server := server.New(server.Config{
		Logger: cfg.Logger,
		Decks:  decks,
		Listen: cfg.Listen,
	})

	serverErrCh := make(chan error, 1)
	go server.Run(serverErrCh)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-monitorErrCh:
		cfg.Logger.Printf("error: (monitor) %v\n", err)
		return err
	case err := <-serverErrCh:
		cfg.Logger.Printf("error: (server) %v\n", err)
		return err
	case <-c:
		cfg.Logger.Println("shutting down...")
	}

	return nil
}
