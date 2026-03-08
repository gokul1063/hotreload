package main

import (
	"fmt"
	"time"

	"hotreload/internal/cli"
	"hotreload/internal/debounce"
	"hotreload/internal/logging"
	"hotreload/internal/watcher"
)

func main() {

	err := logging.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg, err := cli.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	w, err := watcher.New(cfg.Root)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Start()

	db := debounce.New(300 * time.Millisecond)

	fmt.Println("watching:", cfg.Root)

	go func() {
		for range w.Events() {
			db.Trigger()
		}
	}()

	for range db.Events() {
		fmt.Println("debounced change detected")
	}

	fmt.Println("debounced change detected")
}
