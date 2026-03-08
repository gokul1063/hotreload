package main

import (
	"fmt"
	"time"

	"hotreload/internal/builder"
	"hotreload/internal/cli"
	"hotreload/internal/logging"
	"hotreload/internal/runner"
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

	fmt.Println("Root :", cfg.Root)
	fmt.Println("Build:", cfg.Build)
	fmt.Println("Exec :", cfg.Exec)

	b := builder.New(cfg.Build)

	err = b.Build()
	if err != nil {

		fmt.Println("build failed")
		return
	}

	fmt.Println("build success")

	r := runner.New(cfg.Exec)

	err = r.Start()
	if err != nil {

		fmt.Println("server start failed")
		return
	}

	fmt.Println("server running for 10 seconds")

	time.Sleep(10 * time.Second)

	err = r.Stop()
	if err != nil {

		fmt.Println("server stop failed")
		return
	}

	fmt.Println("server stopped")
}
