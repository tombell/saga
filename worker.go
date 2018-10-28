package saga

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/tombell/saga/decks"
)

func worker(watcher *fsnotify.Watcher, d *decks.Decks) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Write != fsnotify.Write {
				return
			}

			fmt.Printf("Reading %s...\n", event.Name)

			snapshot, err := decks.NewSessionSnapshot(event.Name)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}

			tracks := snapshot.NewOrUpdatedTracks(d.Snapshot)

			if err := d.Notify(tracks); err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}

			d.Snapshot = snapshot

			fmt.Println(d)
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}

			fmt.Printf("Error: %v\n", err)
		}
	}
}
