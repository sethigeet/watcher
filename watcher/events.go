package watcher

import (
	"fmt"
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

			eventsChan <- ev.String()
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
		<-eventsChan

		if *cmd.Config.Clear {
			fmt.Printf("\x1bc")
		}

		log.Notice("Refreshing...")

		time.Sleep(*cmd.Config.RunDelay)

		flushEvents()

		run()
	}
}

func flushEvents() {
	for {
		select {
		case <-eventsChan:
			continue
		default:
			return
		}
	}
}
