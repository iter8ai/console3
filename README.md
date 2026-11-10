# console3

A command line tool.

## Usage

```
go build -o bin/console3 ./cmd/console3
./bin/console3 --help
```

## Configuration

Read from environment variables at start-up; see `.env.example`.

| Variable | Required | Default | Flag |
| --- | --- | --- | --- |
| `CONSOLE3_VERBOSE` | no | `false` | `--verbose` |
