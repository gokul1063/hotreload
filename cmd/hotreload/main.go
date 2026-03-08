package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"hotreload/internal/builder"
	"hotreload/internal/cli"
	"hotreload/internal/engine"
	"hotreload/internal/logging"
	"hotreload/internal/runner"
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

		logging.LogError("cli", "Parse", err)
		fmt.Println(err)
		return
	}

	w, err := watcher.New(cfg.Root)
	if err != nil {

		logging.LogError("watcher", "New", err)
		fmt.Println(err)
		return
	}

	b := builder.New(cfg.Build)
	r := runner.New(cfg.Exec)

	e := engine.New(w, b, r)

	// handle ctrl+c / termination
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig
		r.Stop()
		os.Exit(0)
	}()

	err = e.Start()
	if err != nil {
		fmt.Println("engine failed:", err)
	}
}
