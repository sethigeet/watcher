package watcher

import (
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/sethigeet/watcher/watcher/cmd"
)

func sendEvents(w *fsnotify.Watcher) {
	for {
		select {
		case ev, ok := <-w.Events:
			if !ok {
				return
			}

			// HACK: A way around a particular way of saving files(ie. rename original file and write a new file with the new contents)
			if ev.Op == fsnotify.Remove {
				w.Remove(ev.Name)
				exists, err := cmd.Exists(ev.Name)
				if err == nil && exists {
					w.Add(ev.Name)
				}
			}

			eventsChannel <- ev.String()
		case err, ok := <-w.Errors:
			if !ok {
				return
			}
			log.Errorf("An error occurred!\nerror: %s\n", err)
		}
	}
}

func handleEvents() {
	for {
		<-eventsChannel

		time.Sleep(cmd.Config.RunDelay)

		flushEvents()

		log.Notice("Refreshing...")
		errorMessage, ok := run()
		if !ok {
			log.Errorf("Failed running: %s\n", errorMessage)
		}
	}
}

func flushEvents() {
	for {
		select {
		case <-eventsChannel:
			continue
		default:
			return
		}
	}
}
