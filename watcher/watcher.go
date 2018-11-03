package watcher

import (
	"github.com/fsnotify/fsnotify"
)

// WaitForSession ...
func WaitForSession(dir string) (string, error) {
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
