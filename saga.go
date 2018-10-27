package saga

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"

	"github.com/tombell/saga/decks"
	"github.com/tombell/saga/serato"
)

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

	decks := decks.NewDecks()
	decks.Notify(snapshot.Tracks())
	decks.Snapshot = snapshot

	fmt.Println(decks)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go worker(watcher, decks)

	if err := watcher.Add(filepath); err != nil {
		return err
	}

	<-c

	fmt.Println("Shutting down...")

	return nil
}

func worker(watcher *fsnotify.Watcher, decks *decks.Decks) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Write != fsnotify.Write {
				return
			}

			snapshot, err := read(event.Name)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}

			tracks := snapshot.NewOrUpdatedTracks(decks.Snapshot)

			if err := decks.Notify(tracks); err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}

			decks.Snapshot = snapshot

			fmt.Println(decks)
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}

			fmt.Printf("Error: %v\n", err)
		}
	}
}

func read(filepath string) (*decks.SessionSnapshot, error) {
	fmt.Printf("Reading %s...\n", filepath)

	session, err := serato.ReadSession(filepath)
	if err != nil {
		return nil, err
	}

	return decks.NewSessionSnapshot(session), nil
}
