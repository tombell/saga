package saga

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"

	"github.com/tombell/saga/decks"
	"github.com/tombell/saga/web"
)

// Config ...
type Config struct{}

// Run ...
func Run(filepath string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	fmt.Printf("Reading %s...\n", filepath)

	snapshot, err := decks.NewSessionSnapshot(filepath)
	if err != nil {
		return err
	}

	// TODO: Update Notify to take a snapshot, and update the internal snapshot
	// itself.
	d := decks.NewDecks()
	d.Notify(snapshot.Tracks())
	d.Snapshot = snapshot

	fmt.Println(d)

	go worker(watcher, d)

	if err := watcher.Add(filepath); err != nil {
		return err
	}

	serverErrCh := make(chan error, 1)
	server := web.NewServer(d)

	go server.Run(":8080", serverErrCh)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrCh:
		fmt.Printf("Error: (server) %v\n", err)
		return err
	case <-c:
		fmt.Println("Shutting down...")
	}

	return nil
}
