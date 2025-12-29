# Go + Gin Backend Sprint

This repository is a **daily-practice backend sprint** for becoming a Golang + Gin backend developer.

The focus is not on copying code, but on:
- understanding core Go concepts
- designing code that I can write **by myself**
- building backend habits step by step

---

## Goals

- Build strong fundamentals in Go
- Learn Gin by implementing real REST APIs
- Practice clean project structure used in real-world backends
- Commit daily and grow consistently

---

## Tech Stack

- Language: Go
- Web Framework: Gin (from Day 8)
- Environment: WSL2 (Ubuntu) + VS Code
- Version Control: Git / GitHub

---

## Project Structure

```bash

go-gin-sprint/
├── cmd/            # Application entry points
│   └── day01/
│       └── main.go
├── internal/       # Internal application logic
│   └── day01/
│       └── greet.go
├── go.mod
├── go.sum
└── README.md

```

- `cmd/`: executable entry points
- `internal/`: business logic (not meant to be imported externally)

---

## Daily Progress

### Day 1 – Go Basics & CLI Program
**What I learned**
- Initializing a Go module with `go mod`
- Basic project structure (`cmd`, `internal`)
- Handling command-line arguments with `os.Args`
- Separating logic into functions for better testability

**What I built**
- A simple CLI program that prints a greeting
- Default value handling inside a function (not in `main`)

**How to run**
```bash
go run ./cmd/day01
go run ./cmd/day01 Alice
````

---

## Commit Convention

Daily commits follow this format:

```bash
dayXX: short description of what was implemented
```

Examples:

* `day01: implement greet cli with argument handling`
* `day02: handle multiple args with slice iteration`

---

## Roadmap (Planned)

* [x] Day 1: Go basics, CLI, functions
* [ ] Day 2–7: Go fundamentals (structs, interfaces, error handling)
* [ ] Day 8+: Gin REST API, middleware, testing
* [ ] Final: Mini backend project with clean architecture

---

## Notes

This repository is intentionally built step by step.
Each day focuses on writing code that I can reproduce and explain,
not just code that "works".





