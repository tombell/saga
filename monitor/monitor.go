package monitor

import (
	"log"

	"github.com/fsnotify/fsnotify"

	"github.com/tombell/saga/decks"
)

// Config contains configuration and data needed to run the monitor.
type Config struct {
	Logger   *log.Logger
	Decks    *decks.Decks
	Filepath string
}

// Monitor monitors the given file for changes, and updates the status of the
// decks.
type Monitor struct {
	logger   *log.Logger
	decks    *decks.Decks
	filepath string
	watcher  *fsnotify.Watcher
}

// Run begins monitoring the session file for changes, and notifies the decks
// when changes occur.
func (m *Monitor) Run(ch chan error) {
	if err := m.notify(); err != nil {
		m.logger.Printf("error: %v\n", err)
		ch <- err
		return
	}

	for {
		select {
		case event, ok := <-m.watcher.Events:
			if !ok || event.Op&fsnotify.Write != fsnotify.Write {
				continue
			}

			if err := m.notify(); err != nil {
				m.logger.Printf("error: %v\n", err)
				continue
			}
		case err, ok := <-m.watcher.Errors:
			if !ok {
				continue
			}

			m.logger.Printf("error: %v\n", err)
		}
	}
}

// Close closes the file watcher when it's finished being used.
func (m *Monitor) Close() {
	m.watcher.Close()
}

func (m *Monitor) notify() error {
	m.logger.Printf("reading %s...\n", m.filepath)

	snapshot, err := decks.NewSessionSnapshot(m.filepath)
	if err != nil {
		return err
	}

	if err := m.decks.Notify(snapshot); err != nil {
		return err
	}

	for _, deck := range m.decks.Statuses() {
		m.logger.Println(deck)
	}

	return nil
}

// New returns an initialised Monitor ready to watch the session file for
// changes.
func New(cfg Config) (*Monitor, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	m := &Monitor{
		logger:   cfg.Logger,
		decks:    cfg.Decks,
		filepath: cfg.Filepath,
		watcher:  watcher,
	}

	if err := m.watcher.Add(m.filepath); err != nil {
		return nil, err
	}

	return m, nil
}
