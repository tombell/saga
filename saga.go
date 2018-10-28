package saga

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"

	"github.com/tombell/saga/decks"
	"github.com/tombell/saga/serato"
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

	snapshot, err := read(filepath)
	if err != nil {
		return err
	}

	// TODO: Update Notify to take a snapshot, and update the internal snapshot
	// itself.
	decks := decks.NewDecks()
	decks.Notify(snapshot.Tracks())
	decks.Snapshot = snapshot

	fmt.Println(decks)

	go worker(watcher, decks)

	if err := watcher.Add(filepath); err != nil {
		return err
	}

	serverErrCh := make(chan error, 1)
	server := web.NewServer(decks)

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

func read(filepath string) (*decks.SessionSnapshot, error) {
	fmt.Printf("Reading %s...\n", filepath)

	// TODO: move serato.ReadSession into decks.NewSessionSnapshot
	session, err := serato.ReadSession(filepath)
	if err != nil {
		return nil, err
	}

	return decks.NewSessionSnapshot(session), nil
}
