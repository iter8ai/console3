# console3

This project is managed in Iter8, which holds its description and backlog.

## Layout

- `cmd/console3/` — `main`, a thin wrapper around `internal/cli`.
- `internal/cli/` — command line parsing and execution.
- `internal/config/` — configuration read from the environment.

## Run

`go run ./cmd/console3 --help`

## Test

`go test -race ./...`

## Lint

- `gofmt -l .` must print nothing.
- `golangci-lint run` (v2, configured in `.golangci.yml`).

## Build

`go build -o bin/console3 ./cmd/console3`, then `./bin/console3 --help` must exit 0.

## Configuration

All configuration comes from environment variables, read once at start-up into `config.Config`. Nothing is baked in at build time. Every flag has an environment fallback; a flag wins over its variable.

- `CONSOLE3_VERBOSE` (optional, default false; flag `--verbose`)

`.env.example` lists every variable. Never commit `.env`.

## run_checks

In order: `gofmt -l .`, `golangci-lint run`, `go test -race ./...`, `go build -o bin/console3 ./cmd/console3`.
