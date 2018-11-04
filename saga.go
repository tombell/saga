package saga

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"

	"github.com/tombell/saga/api"
	"github.com/tombell/saga/decks"
	"github.com/tombell/saga/monitor"
)

// Config is a data structure for the configuration options passed in from the
// entry point.
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

		file, err := waitForNewSession(cfg.SessionDir)
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

	server := api.New(api.Config{
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

func waitForNewSession(dir string) (string, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return "", err
	}
	defer watcher.Close()

	if err := watcher.Add(dir); err != nil {
		return "", err
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok || event.Op&fsnotify.Create != fsnotify.Create {
				continue
			}

			return event.Name, nil
		case err, ok := <-watcher.Errors:
			if !ok {
				continue
			}

			return "", err
		}
	}
}
