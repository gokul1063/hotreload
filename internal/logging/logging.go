package logging

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

var errorLogger *slog.Logger
var workflowLogger *slog.Logger

var iteration int

func Init() error {

	logDir := ".hotreload/logs"

	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		return err
	}

	err = initIteration(logDir)
	if err != nil {
		return err
	}

	err = initErrorLogger(logDir)
	if err != nil {
		return err
	}

	err = initWorkflowLogger(logDir)
	if err != nil {
		return err
	}

	return nil
}

func initIteration(logDir string) error {

	files, err := os.ReadDir(logDir)
	if err != nil {
		return err
	}

	maxIter := 0

	for _, f := range files {

		var n int

		_, err := fmt.Sscanf(f.Name(), "workflow_%04d.log", &n)
		if err == nil {
			if n > maxIter {
				maxIter = n
			}
		}
	}

	iteration = maxIter + 1

	return nil
}

func initErrorLogger(logDir string) error {

	errorFile := filepath.Join(logDir, "error.log")

	file, err := os.OpenFile(
		errorFile,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)

	if err != nil {
		return err
	}

	handler := slog.NewTextHandler(file, &slog.HandlerOptions{
		AddSource: false,
	})

	errorLogger = slog.New(handler)

	return nil
}

func initWorkflowLogger(logDir string) error {

	name := fmt.Sprintf("workflow_%04d.log", iteration)

	path := filepath.Join(logDir, name)

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	handler := slog.NewTextHandler(file, &slog.HandlerOptions{
		AddSource: false,
	})

	workflowLogger = slog.New(handler)

	return nil
}

func LogError(pkg string, fn string, err error) {

	if err == nil {
		return
	}

	errorLogger.Error(
		err.Error(),
		"package", pkg,
		"function", fn,
		"time", time.Now().UTC(),
	)
}

func LogWorkflow(pkg string, fn string, status string) {

	workflowLogger.Info(
		"workflow",
		"package", pkg,
		"function", fn,
		"status", status,
		"time", time.Now().UTC(),
	)
}
