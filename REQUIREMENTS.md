# Functional Requirements

FR1
The tool must monitor a project directory for file changes.

FR2
The tool must support recursive directory watching.

FR3
When a file change occurs, the tool must trigger a rebuild.

FR4
After a successful build, the tool must restart the server.

FR5
The tool must execute the build command provided through CLI arguments.

FR6
The tool must execute the run command provided through CLI arguments.

FR7
The first build must run immediately when the program starts.

FR8
The server logs must stream in real time.

FR9
The tool must terminate the previous server process before starting a new one.

FR10
The tool must kill the full process group of the server.

FR11
The tool must debounce rapid file system events.

FR12
The tool must detect newly created directories and begin watching them.

FR13
The tool must handle deleted directories without crashing.

FR14
The tool must ignore irrelevant files and directories.

FR15
The tool must log all errors returned by functions.

FR16
The tool must generate a workflow log for each execution of the program.

---

# Non Functional Requirements

NFR1
File change detection should trigger rebuild within approximately 2 seconds.

NFR2
The system must run continuously for long durations.

NFR3
The system must avoid excessive rebuild loops.

NFR4
The system must manage operating system resources safely.

NFR5
Logging must not block the execution pipeline.

---

# Logging Requirements

Error Log

- append-only
- persistent across runs
- records all function errors
- stored in `.hotreload/logs/error.log`

Workflow Log

- created once per application run
- stored in `.hotreload/logs/workflow_<iteration>.log`
- records function calls and status

---

# Platform Requirements

Operating System

Linux

Process Control

Must support process group termination using POSIX signals.

---

# Allowed Libraries

Allowed

fsnotify

Forbidden

air
reflex
realize
