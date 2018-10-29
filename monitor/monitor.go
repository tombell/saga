package monitor

import (
	"log"

	"github.com/fsnotify/fsnotify"

	"github.com/tombell/saga/decks"
)

// Config contains configuration and data needed to run the monitor.
type Config struct {
	Logger   *log.Logger
	Filepath string
	Decks    *decks.Decks
}

// Monitor monitors the given file for changes, and updates the status of the
// decks.
type Monitor struct {
	logger   *log.Logger
	filepath string
	decks    *decks.Decks
	watcher  *fsnotify.Watcher
}

// Run ...
func (m *Monitor) Run(ch chan error) {
	for {
		select {
		case event, ok := <-m.watcher.Events:
			if !ok || event.Op&fsnotify.Write != fsnotify.Write {
				return
			}

			m.logger.Printf("reading %s...\n", event.Name)

			snapshot, err := decks.NewSessionSnapshot(event.Name)
			if err != nil {
				m.logger.Printf("error: %v\n", err)
				return
			}

			if err := m.decks.Notify(snapshot); err != nil {
				m.logger.Printf("error: %v\n", err)
				return
			}

			m.logger.Println(m.decks)
		case err, ok := <-m.watcher.Errors:
			if !ok {
				return
			}

			m.logger.Printf("error: %v\n", err)
		}
	}
}

// Close closes the file watcher when it's finished being used.
func (m *Monitor) Close() {
	m.watcher.Close()
}

// New ...
func New(cfg Config) (*Monitor, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	m := &Monitor{
		logger:   cfg.Logger,
		filepath: cfg.Filepath,
		decks:    cfg.Decks,
		watcher:  watcher,
	}

	if err := m.watcher.Add(m.filepath); err != nil {
		return nil, err
	}

	return m, nil
}
