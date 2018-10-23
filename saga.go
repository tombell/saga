package saga

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/tombell/saga/serato"
)

// Run ...
func Run(filepath string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	done := make(chan bool)
	go worker(watcher)

	if err := watcher.Add(filepath); err != nil {
		return err
	}

	if err := read(filepath); err != nil {
		return err
	}

	<-done

	fmt.Println("Shutting down...")

	return nil
}

func worker(watcher *fsnotify.Watcher) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Write != fsnotify.Write {
				return
			}

			if err := read(event.Name); err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}

			fmt.Printf("Error: %v\n", err)
		}
	}
}

func read(filepath string) error {
	fmt.Printf("Reading %s...\n", filepath)

	session, err := serato.ReadSession(filepath)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("VRSN: %s\n", session.Vrsn.Version())
	fmt.Printf("OENT: %d\n", len(session.Oent))
	fmt.Printf("OREN: %d\n", len(session.Oren))
	fmt.Println()

	for _, oent := range session.Oent {
		fmt.Printf("%-3v: Deck %v (A:%-5v / P:%-5v) [%-7v] %s - %s\n",
			oent.Adat.Row,
			oent.Adat.Deck,
			oent.Adat.Added,
			oent.Adat.Played,
			oent.Adat.Status(),
			oent.Adat.Artist,
			oent.Adat.Title,
		)
	}

	fmt.Println()

	return nil
}
