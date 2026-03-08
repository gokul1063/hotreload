package main

import (
	"fmt"

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

	logging.LogWorkflow("main", "ParseCLI", "success")

	w, err := watcher.New(cfg.Root)
	if err != nil {

		logging.LogError("watcher", "New", err)
		fmt.Println(err)
		return
	}

	b := builder.New(cfg.Build)

	r := runner.New(cfg.Exec)

	e := engine.New(w, b, r)

	err = e.Start()
	if err != nil {

		fmt.Println("engine failed:", err)
	}
}
