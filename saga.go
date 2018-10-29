package saga

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/tombell/saga/decks"
	"github.com/tombell/saga/monitor"
	"github.com/tombell/saga/server"
)

// Config ...
type Config struct {
	Logger   *log.Logger
	Listen   string
	Filepath string
}

// Run begins the process of listening for changes to the given Serato session
// file. It keeps a realtime model of the decks, and the tracks they're
// playing, and have played.
func Run(cfg Config) error {
	cfg.Logger.Printf("reading %s\n", cfg.Filepath)

	snapshot, err := decks.NewSessionSnapshot(cfg.Filepath)
	if err != nil {
		return err
	}

	decks := decks.NewDecks(cfg.Logger)

	if err := decks.Notify(snapshot); err != nil {
		return err
	}

	// TODO: make nicer, to return slice of statuses
	for _, deck := range strings.Split(decks.String(), "\n") {
		cfg.Logger.Println(deck)
	}

	monitor, err := monitor.New(monitor.Config{
		Logger:   cfg.Logger,
		Decks:    decks,
		Filepath: cfg.Filepath,
	})
	if err != nil {
		return err
	}
	defer monitor.Close()

	monitorErrCh := make(chan error, 1)
	go monitor.Run(monitorErrCh)

	server := server.New(server.Config{
		Logger:  cfg.Logger,
		Decks:   decks,
		Address: cfg.Listen,
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
