package watcher

import (
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"

	"hotreload/internal/filter"
	"hotreload/internal/logging"
)

type Watcher struct {
	root    string
	watcher *fsnotify.Watcher
	events  chan struct{}
}

func New(root string) (*Watcher, error) {

	w, err := fsnotify.NewWatcher()
	if err != nil {
		logging.LogError("watcher", "New", err)
		return nil, err
	}

	obj := &Watcher{
		root:    root,
		watcher: w,
		events:  make(chan struct{}, 1),
	}

	err = obj.addRecursive(root)
	if err != nil {
		logging.LogError("watcher", "addRecursive", err)
		return nil, err
	}

	logging.LogWorkflow("watcher", "New", "success")

	return obj, nil
}

func (w *Watcher) Events() <-chan struct{} {
	return w.events
}

func (w *Watcher) Start() {

	logging.LogWorkflow("watcher", "Start", "started")

	go func() {

		for {
			select {

			case event, ok := <-w.watcher.Events:

				if !ok {
					return
				}

				w.handleEvent(event)

			case err, ok := <-w.watcher.Errors:

				if !ok {
					return
				}

				logging.LogError("watcher", "Events", err)
			}
		}
	}()
}

func (w *Watcher) handleEvent(event fsnotify.Event) {

	path := event.Name

	if filter.ShouldIgnore(path) {
		return
	}

	// if new directory created -> watch it
	if event.Op&fsnotify.Create == fsnotify.Create {

		info, err := os.Stat(path)
		if err == nil && info.IsDir() {

			err := w.addRecursive(path)
			if err != nil {
				logging.LogError("watcher", "addRecursive", err)
			}
		}
	}

	select {
	case w.events <- struct{}{}:
	default:
	}

	logging.LogWorkflow("watcher", "handleEvent", "triggered")
}

func (w *Watcher) addRecursive(dir string) error {

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if filter.ShouldIgnore(path) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if d.IsDir() {

			err := w.watcher.Add(path)
			if err != nil {
				return err
			}

			logging.LogWorkflow("watcher", "AddDir", path)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (w *Watcher) Close() error {

	logging.LogWorkflow("watcher", "Close", "called")

	return w.watcher.Close()
}
