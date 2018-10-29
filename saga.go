package saga

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"

	"github.com/tombell/saga/decks"
	"github.com/tombell/saga/web"
)

// Config ...
type Config struct {
	Filepath string
	Logger   *log.Logger
}

// Run begins the process of listening for changes to the given Serato session
// file. It keeps a realtime model of the decks, and the tracks they're
// playing, and have played.
func Run(cfg Config) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	cfg.Logger.Printf("reading %s\n", cfg.Filepath)

	snapshot, err := decks.NewSessionSnapshot(cfg.Filepath)
	if err != nil {
		return err
	}

	d := decks.NewDecks(cfg.Logger)

	if err := d.Notify(snapshot); err != nil {
		return err
	}

	cfg.Logger.Println(d)

	go worker(watcher, d)

	if err := watcher.Add(cfg.Filepath); err != nil {
		return err
	}

	serverErrCh := make(chan error, 1)
	server := web.NewServer(d)

	go server.Run(":8080", serverErrCh)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrCh:
		cfg.Logger.Printf("Error: (server) %v\n", err)
		return err
	case <-c:
		cfg.Logger.Println("Shutting down...")
	}

	return nil
}
