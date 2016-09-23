package main

import (
	"log"

	"github.com/dietsche/rfsnotify"
)

func main() {
	watcher, err := rfsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Events:
				log.Println("event:", ev)
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.AddRecursive("/tmp/")
	if err != nil {
		log.Fatal(err)
	}

	// Hang so program doesn't exit
	<-done
	
	/* ... do stuff ... */
	watcher.Close()
}
