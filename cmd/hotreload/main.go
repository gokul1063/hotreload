package main

import (
	"fmt"

	"hotreload/internal/cli"
	"hotreload/internal/filter"
	"hotreload/internal/logging"
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
	fmt.Println(filter.ShouldIgnore(".git"))
	fmt.Println(filter.ShouldIgnore("main.go"))
	fmt.Println(filter.ShouldIgnore("file.go~"))
}
