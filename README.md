# hotreload

hotreload is a CLI tool that automatically rebuilds and restarts a server when source files change.
It eliminates the need to manually stop, rebuild, and restart applications during development.

The tool watches a project directory, detects file modifications, rebuilds the project, and restarts the server process.

This project was developed as a backend engineering assignment and focuses on correctness, stability, and system-level robustness.

---

# Features

Core Features

- Watches a project directory recursively
- Detects file changes
- Automatically rebuilds the project
- Automatically restarts the server
- Streams server logs in real time
- Performs an initial build immediately when the tool starts

Process Handling

- Ensures previous server process is fully terminated before restart
- Kills the entire process group, not just the parent process

Performance

- Debounces rapid file events
- Avoids redundant rebuilds caused by editor save patterns

Project Structure Handling

- Watches nested folders
- Detects new folders created while running
- Handles folder deletion gracefully

File Filtering

The watcher ignores unnecessary files and directories:

.git
node_modules
build artifacts
temporary editor files
linux temporary files

Logging System

Two logging systems are implemented.

Error Log

Persistent log that records all function errors.

Location:

```
.hotreload/logs/error.log
```

Format:

```
timestamp | ERROR | package.function | message
```

Example:

```
2026-03-08T17:10:23Z | ERROR | builder.Build | go build failed: exit status 1
```

Workflow Log

A temporary debug log that records the execution flow of the program.

A **new workflow log is created every time the application starts**.

Example:

```
.hotreload/logs/workflow_0001.log
.hotreload/logs/workflow_0002.log
```

Format:

```
timestamp | package.function | status
```

Example:

```
2026-03-08T17:10:23Z | engine.Start | started
2026-03-08T17:10:23Z | watcher.Init | success
2026-03-08T17:10:24Z | builder.Build | started
2026-03-08T17:10:25Z | builder.Build | success
```

Workflow logs are intended for debugging and may be removed in production builds.

---

# CLI Usage

```
hotreload \
--root ./myproject \
--build "go build -o ./bin/server ./cmd/server" \
--exec "./bin/server"
```

Parameters

--root
Directory to watch for file changes.

--build
Command used to build the project.

--exec
Command used to run the server.

---

# Example

Example project structure:

```
project/
    cmd/server/main.go
    internal/
    bin/
```

Run hotreload:

```
hotreload \
--root ./project \
--build "go build -o ./bin/server ./cmd/server" \
--exec "./bin/server"
```

When a `.go` file is saved:

1. change is detected
2. build command executes
3. previous server process is terminated
4. new server process starts

---

# Repository Structure

```
hotreload
│
├── cmd/
│   └── hotreload/
│       └── main.go
│
├── internal/
│   ├── engine/
│   ├── watcher/
│   ├── builder/
│   ├── runner/
│   ├── debounce/
│   ├── filter/
│   └── logging/
│
├── testserver/
│   └── main.go
│
├── scripts/
│   └── demo.sh
│
├── README.md
├── REQUIREMENTS.md
└── TODO.md
```

---

# Demo

A sample server is provided in `testserver/`.

To test hot reload:

```
cd testserver
go run main.go
```

Then run hotreload pointing to that folder.

---

# Dependencies

Allowed dependency:

```
github.com/fsnotify/fsnotify
```

All other functionality is implemented manually.

---

# Development Notes

Key design goals:

- stable long-running process
- predictable process control
- minimal rebuild latency
- clear logging
- clean modular architecture
