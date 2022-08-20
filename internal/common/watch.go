package common

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/urfave/cli/v2"
)

func Watch(cCtx *cli.Context) error {
	// Create a new watcher.
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer w.Close()

	files := cCtx.StringSlice("file")

	// Start listening for events.
	go fileLoop(w, files, cCtx)

	// Add all files.
	for _, p := range files {
		st, err := os.Lstat(p)
		if err != nil {
			return err
		}

		if st.IsDir() {
			return fmt.Errorf("%s is a directory, not a file", p)
		}

		// Watch the directory, not the file itself.
		err = w.Add(filepath.Dir(p))
		if err != nil {
			return fmt.Errorf("%q: %s", p, err)
		}
	}

	log.Println("Watching for changes: press ^C to exit")
	<-make(chan struct{}) // Block forever
	return nil
}

func fileLoop(w *fsnotify.Watcher, files []string, cCtx *cli.Context) {
	for {
		select {
		// Read from Errors.
		case err, ok := <-w.Errors:
			if !ok { // Channel was closed (i.e. Watcher.Close() was called).
				return
			}
			log.Printf("ERROR: %s\n", err)
		// Read from Events.
		case e, ok := <-w.Events:
			if !ok { // Channel was closed (i.e. Watcher.Close() was called).
				return
			}

			var found bool

			for _, f := range files {
				if f == e.Name {
					found = true
				}
			}

			if !found {
				continue
			}

			if e.Op != fsnotify.Write {
				continue
			}

			log.Printf("%s changed reloading\n", e.Name)

			err := Exec(cCtx)
			if err != nil {
				log.Printf("ERROR: %s\n", err)
			}
		}
	}
}
