# Phase 1 — Repository Setup

- initialize git repository
- create Go module
- create folder structure
- add fsnotify dependency
- create logging directory generator

---

# Phase 2 — Logging System

Implement logging package

Tasks

- create error logger
- create workflow logger
- implement iteration detection
- create workflow log file per run
- implement logging helpers

Functions

logging.Init()

logging.LogError(pkg, func, err)

logging.LogWorkflow(pkg, func, status)

---

# Phase 3 — CLI Parsing

Tasks

- parse CLI arguments
- validate parameters
- ensure root directory exists
- validate build command
- validate exec command

Functions

cli.ParseArgs()

---

# Phase 4 — File Watcher

Tasks

- initialize fsnotify watcher
- recursively watch directories
- detect new directories
- detect directory deletion
- filter ignored files

Functions

watcher.Init()

watcher.AddRecursive()

watcher.HandleEvents()

---

# Phase 5 — Debounce System

Tasks

- implement change aggregation
- prevent rapid rebuild loops
- collapse multiple file events

Functions

debounce.New()

debounce.Trigger()

---

# Phase 6 — Build System

Tasks

- execute build command
- capture stdout and stderr
- return build status
- propagate errors to logger

Functions

builder.Build()

---

# Phase 7 — Process Runner

Tasks

- start server process
- attach stdout and stderr
- create process group
- kill previous server
- ensure full process tree termination

Functions

runner.Start()

runner.Stop()

runner.KillGroup()

---

# Phase 8 — Engine

Tasks

- coordinate watcher, builder, runner
- manage rebuild pipeline
- handle build cancellation
- prevent overlapping builds

Functions

engine.Start()

engine.HandleChange()

engine.Restart()

---

# Phase 9 — Test Server

Tasks

- create simple HTTP server
- log requests
- demonstrate restart behavior

---

# Phase 10 — Demo Script

Tasks

- add Makefile or script
- automate demo steps
- prepare Loom recording
