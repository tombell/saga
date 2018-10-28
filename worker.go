package saga

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/tombell/saga/decks"
)

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
