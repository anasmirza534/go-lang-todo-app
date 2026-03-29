# go-lang-todo-app

A Todo app built with Go — learning project.

## Stack

- Go 1.21
- SQLite
- REST API + CLI + HTMX web UI (planned)

## Getting Started

```bash
go mod download
go run main.go
```

## Migration

install goose

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

create

```bash
goose create [migration-name] sql
```

status

```bash
goose status
```

up

```bash
goose up
```

down

```bash
goose down
```

## License

MIT
