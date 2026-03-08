package engine

import (
	"time"

	"hotreload/internal/builder"
	"hotreload/internal/debounce"
	"hotreload/internal/logging"
	"hotreload/internal/runner"
	"hotreload/internal/watcher"
)

type Engine struct {
	watcher *watcher.Watcher
	builder *builder.Builder
	runner  *runner.Runner
	deb     *debounce.Debouncer
}

func New(
	w *watcher.Watcher,
	b *builder.Builder,
	r *runner.Runner,
) *Engine {

	return &Engine{
		watcher: w,
		builder: b,
		runner:  r,
		deb:     debounce.New(300 * time.Millisecond),
	}
}

func (e *Engine) Start() error {

	logging.LogWorkflow("engine", "Start", "started")

	err := e.builder.Build()
	if err != nil {

		logging.LogError("engine", "InitialBuild", err)
		return err
	}

	err = e.runner.Start()
	if err != nil {

		logging.LogError("engine", "RunnerStart", err)
		return err
	}

	e.watcher.Start()

	go func() {
		for range e.watcher.Events() {
			e.deb.Trigger()
		}
	}()

	go func() {
		for range e.deb.Events() {
			e.reload()
		}
	}()

	logging.LogWorkflow("engine", "Start", "running")

	select {}
}

func (e *Engine) reload() {

	logging.LogWorkflow("engine", "Reload", "triggered")

	err := e.builder.Build()
	if err != nil {

		logging.LogError("engine", "Build", err)
		return
	}

	err = e.runner.Stop()
	if err != nil {

		logging.LogError("engine", "RunnerStop", err)
	}

	err = e.runner.Start()
	if err != nil {

		logging.LogError("engine", "RunnerStart", err)
	}
}
