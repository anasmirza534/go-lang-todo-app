# Go Learning Checklist — For Experienced Backend Devs

> **Audience:** Backend developers coming from PHP, Node.js/TypeScript, or similar. Target project: Todo app with REST API, CLI, HTMX web UI, SQLite + migrations — plus auth and OpenAPI docs.
>
> This checklist skips web fundamentals you already know. Focus is on what's _different_ or _new_ in Go.

---

> **Sharing this doc?** It's written for devs who already know backend concepts (REST, HTTP, SQL, auth). The final section has a PHP/TS → Go equivalents table as a quick reference. If you're coming from a different language, the concepts still apply — just adapt the comparisons.

---

## 1. The Go Mental Model (Read This First)

- [ ] **No classes, no inheritance** — only structs + interfaces (composition-first)
- [ ] **No exceptions** — errors are _values_, returned explicitly (no try/catch)
- [ ] **Compiled, statically typed** — like TS but no runtime, no JIT
- [ ] **Go is opinionated** — one formatter (`gofmt`), one test runner, one build system
- [ ] **Concurrency is a first-class citizen** — goroutines are not threads, channels are not events
- [ ] **Zero values** — every type has a zero value (int=0, string="", bool=false, pointer=nil)
- [ ] **Exported vs unexported** — UPPERCASE = public, lowercase = package-private (no keywords)

---

## 2. Toolchain & Setup

- [ ] `go mod init`, `go.mod`, `go.sum` — like `package.json` but simpler
- [ ] `go run`, `go build`, `go install` — know the difference
- [ ] `go get` — adding dependencies
- [ ] `go fmt` / `gofmt` — auto-formatter, non-negotiable
- [ ] `go vet` — static analysis
- [ ] `go test ./...` — built-in test runner
- [ ] `gopls` — language server for your editor (VS Code: Go extension)
- [ ] Air / `go-watch` — hot reload for development

---

## 3. Core Syntax You'll Hit Immediately

### Types & Variables

- [ ] Short declaration `:=` vs `var` — when to use which
- [ ] Multiple return values — `return result, err` (you'll use this constantly)
- [ ] Named return values (optional, used in some stdlib)
- [ ] Blank identifier `_` — discard values you don't need
- [ ] Type aliases vs type definitions (`type UserID int`)

### Structs (Your New Classes)

- [ ] Struct definition and field access
- [ ] Struct embedding (composition over inheritance)
- [ ] Methods on structs — value receivers vs pointer receivers (important!)
- [ ] Anonymous structs — useful for JSON, test data
- [ ] Struct tags — `json:"name"`, `db:"name"` (crucial for REST + SQL)

### Interfaces

- [ ] Interfaces are **implicit** — no `implements` keyword
- [ ] Small interfaces are idiomatic — `io.Reader`, `io.Writer` are 1-2 methods
- [ ] `interface{}` / `any` — like `unknown` in TS, avoid overusing
- [ ] Type assertions and type switches — `val.(type)`, `switch v := x.(type)`

### Error Handling

- [ ] The `error` interface — `errors.New()`, `fmt.Errorf()`
- [ ] `if err != nil` pattern — you'll write this 100x per day
- [ ] Wrapping errors with `%w` — `fmt.Errorf("context: %w", err)`
- [ ] `errors.Is()` and `errors.As()` — like checking instanceof in JS
- [ ] Custom error types — struct implementing `error` interface
- [ ] **When NOT to use panic** — only for unrecoverable programmer errors

### Slices & Maps

- [ ] Slice internals — backed by array, `len` vs `cap`
- [ ] `append()` — may allocate new array under the hood
- [ ] `make([]T, len, cap)` — pre-allocate when you know size
- [ ] Maps — `make(map[string]T)`, always check key existence with `val, ok := m[key]`
- [ ] Nil slice vs empty slice — subtle difference, matters for JSON (`null` vs `[]`)

### Pointers

- [ ] `&` (address-of) and `*` (dereference)
- [ ] When to use pointers: large structs, mutation, optional values
- [ ] No pointer arithmetic (simpler than C)
- [ ] Go GC handles memory — no `free()`

---

## 4. Packages & Project Structure

- [ ] One package per directory (convention)
- [ ] `main` package + `main()` func = executable
- [ ] Internal packages — `internal/` dir restricts import to parent module
- [ ] Recommended project layout for a small app:

```
todo-app/
├── main.go
├── go.mod
├── cmd/              # CLI entry points (cobra commands)
│   ├── root.go
│   ├── add.go
│   └── list.go
├── internal/
│   ├── db/           # SQLite, migrations
│   ├── handler/      # HTTP handlers
│   ├── model/        # Structs / domain types
│   └── service/      # Business logic
├── web/
│   └── templates/    # HTML templates for HTMX
└── migrations/       # SQL migration files
```

- [ ] Avoid circular imports — Go won't compile with them (forces clean layering)

---

## 5. Concurrency (Go's Superpower — Understand Basics)

> You don't need deep mastery for a todo app, but you'll encounter this in stdlib code.

- [ ] **Goroutines** — `go func()` — lightweight, not OS threads (100k+ is fine)
- [ ] **Channels** — typed message passing between goroutines (`chan T`)
  - [ ] Buffered vs unbuffered channels
  - [ ] Sending `ch <- val` and receiving `val := <-ch`
  - [ ] Closing channels — `close(ch)`, ranging over channel
- [ ] **`select` statement** — like `switch` but for channels (multiplexing)
- [ ] **`sync.WaitGroup`** — wait for goroutines to finish (like `Promise.all`)
- [ ] **`sync.Mutex`** — protect shared state (like a lock)
- [ ] **`context.Context`** — cancellation, timeouts, request-scoped values (you WILL use this in HTTP)
  - [ ] `context.WithTimeout`, `context.WithCancel`
  - [ ] Passing `ctx` as first param to functions — idiomatic
  - [ ] `ctx.Done()` channel — listen for cancellation

> **For your todo app:** You mostly need `context.Context` for HTTP handlers and DB calls. Goroutines/channels are optional at this scale.

---

## 6. Defer, Panic, Recover

- [ ] `defer` — runs at end of function, LIFO order (use for cleanup: close files, DB rows)
- [ ] `defer rows.Close()` — you'll use this pattern with every SQL query
- [ ] `panic` — like throwing an uncatchable exception (rare, avoid in library code)
- [ ] `recover` — catch a panic, only works inside a deferred func (use in HTTP middleware)

---

## 7. Standard Library Highlights

> Go's stdlib is excellent — resist reaching for external packages until you know what's built in.

- [ ] `fmt` — printing, formatting, `Errorf`
- [ ] `net/http` — HTTP server + client (no framework needed for simple APIs)
  - [ ] `http.HandleFunc`, `http.ListenAndServe`
  - [ ] `http.Request`, `http.ResponseWriter`
  - [ ] Reading JSON body, writing JSON response
- [ ] `html/template` — safe HTML templating (auto-escapes, good for HTMX)
- [ ] `encoding/json` — `json.Marshal`, `json.Unmarshal`, `json.NewDecoder`
- [ ] `database/sql` — generic DB interface (works with any driver)
  - [ ] `db.Query`, `db.QueryRow`, `db.Exec`
  - [ ] `rows.Scan` — mapping columns to struct fields
  - [ ] Prepared statements
- [ ] `os` — env vars, file I/O, stdin/stdout
- [ ] `log` / `log/slog` — structured logging (slog is the modern one, Go 1.21+)
- [ ] `strings`, `strconv` — string ops (more explicit than JS)
- [ ] `time` — time handling (very clean API)
- [ ] `flag` — basic CLI flags (use cobra instead for your project)

---

## 8. HTTP & REST API Patterns

- [ ] **Router choice for your project:** [Chi](https://github.com/go-chi/chi) — lightweight, stdlib-compatible, good middleware support
- [ ] Middleware pattern — `func(http.Handler) http.Handler`
- [ ] JSON response helper pattern — write a `writeJSON(w, status, data)` helper
- [ ] Reading + validating request body
- [ ] URL params with Chi — `chi.URLParam(r, "id")`
- [ ] HTTP status codes — `http.StatusOK`, `http.StatusCreated`, etc.
- [ ] CORS middleware — if you add a frontend framework later

---

## 9. SQLite + Migrations

### Driver

- [ ] **Use `modernc.org/sqlite`** (pure Go, no CGO needed) — simpler than `mattn/go-sqlite3`
- [ ] Open DB: `sql.Open("sqlite", "./todos.db")`
- [ ] Connection pool settings for SQLite: `db.SetMaxOpenConns(1)` (SQLite is single-writer)

### Migrations

- [ ] **[goose](https://github.com/pressly/goose)** — simple, SQL-first migrations (like Laravel migrations but SQL files)
- [ ] Migration file naming: `001_create_todos.sql`, `002_add_priority.sql`
- [ ] Up/down migrations
- [ ] Embedding migrations with `//go:embed migrations/*.sql`

### Patterns

- [ ] Always use `context.Context` with DB calls: `db.QueryContext(ctx, ...)`
- [ ] `defer rows.Close()` after every `Query`
- [ ] Scan into struct fields manually (no ORM magic by default)
- [ ] **Optional ORM:** [sqlc](https://sqlc.dev/) — generates Go code from SQL queries (recommended over raw scanning)

---

## 10. HTML Templates + HTMX

- [ ] `html/template` package — `template.ParseFiles`, `template.ParseGlob`
- [ ] Template execution — `tmpl.Execute(w, data)`
- [ ] Passing structs to templates as data
- [ ] Template actions — `{{.FieldName}}`, `{{range .}}`, `{{if .Condition}}`
- [ ] Partial templates — define named blocks for HTMX partial responses
- [ ] Returning partial HTML from handlers (HTMX swaps these in)
- [ ] Template inheritance pattern with `{{block}}` and `{{template}}`

---

## 11. CLI with Cobra

- [ ] `cobra-cli init` — scaffolds the project
- [ ] Root command + subcommands pattern
- [ ] Flags: `cmd.Flags().StringP(...)`, `cmd.Flags().BoolP(...)`
- [ ] Persistent flags — available to all subcommands
- [ ] `RunE` vs `Run` — use `RunE` to return errors properly
- [ ] Sharing DB/config between commands via struct or closure

---

## 12. Testing

- [ ] `_test.go` files — convention, not configuration
- [ ] `func TestXxx(t *testing.T)` — naming matters
- [ ] `t.Error`, `t.Fatal`, `t.Run` (subtests)
- [ ] Table-driven tests — idiomatic Go testing pattern
- [ ] `httptest.NewRecorder()` — test HTTP handlers without a real server
- [ ] `testify` library — `assert.Equal`, `require.NoError` (optional but very helpful)
- [ ] In-memory SQLite for tests — `sql.Open("sqlite", ":memory:")`

---

## 13. Coding Conventions (Go is Opinionated)

- [ ] **Run `gofmt` always** — no debates about style
- [ ] **Short variable names are idiomatic** — `i`, `r`, `w`, `err`, `ctx` are normal
- [ ] **Error variables named `err`** — always, not `error` or `e`
- [ ] **Receiver names** — short, consistent (use `t` for `Todo`, not `this`/`self`)
- [ ] **Avoid stuttering** — don't write `todo.TodoService`, just `todo.Service`
- [ ] **Comments on exported symbols** — start with the name: `// Todo represents a task`
- [ ] **No unused imports or variables** — code won't compile (enforced)
- [ ] **Return early** — avoid deep nesting, handle errors immediately and return
- [ ] **Accept interfaces, return structs** — flexible function signatures
- [ ] Linter: `golangci-lint` — configure once, run in CI

---

## 14. Common Gotchas (PHP/JS Dev Edition)

| Gotcha                | Explanation                                                                                       |
| --------------------- | ------------------------------------------------------------------------------------------------- |
| `nil` vs zero value   | A nil pointer panics, a zero struct doesn't — initialize your structs                             |
| Goroutine closures    | Loop variable capture bug — pass variable as arg, don't close over `i` in goroutines              |
| Map concurrency       | Maps are NOT thread-safe — use `sync.Map` or mutex if shared across goroutines                    |
| JSON field casing     | Struct fields must be exported (UPPERCASE) to be marshalled; use `json:"name"` tag for casing     |
| `range` copies values | `for i, v := range slice` — `v` is a copy; use index `slice[i]` to mutate                         |
| Nil interface         | A nil interface and an interface holding a nil pointer are NOT equal (common panic source)        |
| `defer` in loops      | `defer` inside a loop doesn't run until function ends, not loop iteration — pull into helper func |
| Context propagation   | If you don't thread `ctx` through, you lose cancellation — always accept `ctx` as first arg       |

---

## 15. Your Todo App — Build Order

```
Phase 1: Core
  1. go mod init + project structure
  2. SQLite setup (modernc.org/sqlite) + goose migrations
  3. Todo CRUD: model struct, db layer (raw sql or sqlc)
  4. REST API with Chi: GET/POST/PUT/DELETE /todos

Phase 2: CLI
  5. cobra setup
  6. Commands: add, list, done, delete
  7. CLI reads from same SQLite DB

Phase 3: Web UI
  8. html/template setup
  9. Base layout + todo list template
 10. HTMX: add todo inline, toggle done, delete (partial responses)
 11. Static file serving (CSS)

Phase 4: Polish
 12. Structured logging (slog)
 13. Config via env vars (os.Getenv or godotenv)
 14. Graceful shutdown (signal handling)
 15. Tests for handlers + DB layer

Phase 5: Auth
 16. Password hashing (bcrypt)
 17. Cookie-based sessions for HTMX web (gorilla/sessions or signed cookies)
 18. JWT middleware for REST API
 19. CSRF protection for web forms
 20. Auth middleware — protect routes in Chi

Phase 6: OpenAPI & React-ready API
 21. Add swaggo/swag annotations to handlers
 22. Generate swagger.json / openapi.yaml
 23. Embed Swagger UI for dev access
 24. Share spec with frontend dev (React client)
```

---

## 16. Auth — Cookie Sessions (HTMX Web) + JWT (REST API)

> **Not needed for Phase 1–4. Add after the core app is working.**

### Password Handling

- [ ] `golang.org/x/crypto/bcrypt` — `bcrypt.GenerateFromPassword`, `bcrypt.CompareHashAndPassword`
- [ ] Never store plain passwords — always hash before saving to DB
- [ ] Cost factor: 12 is a safe default

### Cookie-Based Sessions (for HTMX / Server-Rendered UI)

- [ ] **[`gorilla/sessions`](https://github.com/gorilla/sessions)** — cookie or server-side sessions
  - [ ] `sessions.NewCookieStore([]byte(secretKey))` — signed + optionally encrypted
  - [ ] `session.Set("user_id", id)` / `session.Save(r, w)`
  - [ ] `session.Get("user_id")` in middleware
- [ ] **Alternative (simpler):** signed cookie with HMAC — no external lib, stateless
- [ ] Set `HttpOnly`, `Secure`, `SameSite=Lax` flags on cookies — prevents XSS/CSRF
- [ ] **CSRF protection** — required for cookie auth + form submissions
  - [ ] [`gorilla/csrf`](https://github.com/gorilla/csrf) — middleware + template helper
  - [ ] CSRF token in hidden form field or `X-CSRF-Token` header (HTMX can send this)
- [ ] Session-based auth middleware for Chi — check cookie, load user, attach to `ctx`

### JWT (for REST API)

- [ ] **[`golang-jwt/jwt`](https://github.com/golang-jwt/jwt)** — standard JWT library for Go
- [ ] `jwt.NewWithClaims(jwt.SigningMethodHS256, claims)` — generate token
- [ ] `jwt.ParseWithClaims(tokenStr, &claims, keyFunc)` — validate token
- [ ] Custom claims struct — embed `jwt.RegisteredClaims`, add your fields (`UserID`, `Role`)
- [ ] Token expiry — always set `ExpiresAt` in claims
- [ ] Store JWT secret in env var, never hardcode
- [ ] **JWT middleware for Chi:**
  - [ ] Extract `Authorization: Bearer <token>` header
  - [ ] Validate + parse claims
  - [ ] Attach user to `ctx` — `context.WithValue(ctx, userKey, user)`
  - [ ] Return `401` if missing/invalid, `403` if insufficient permissions
- [ ] **Refresh tokens** (optional, do later): short-lived access token + long-lived refresh token stored in DB

### Two Auth Systems, One App

```
/api/*          → JWT middleware (for React client / mobile)
/app/*          → Cookie session middleware (for HTMX web)
/auth/login     → issues both cookie AND returns JWT (or separate endpoints)
```

- [ ] Keep auth logic in `internal/auth/` — shared between both middlewares
- [ ] Password check, user lookup, token generation are reusable regardless of transport

---

## 17. OpenAPI & Swagger Docs (for Frontend Dev)

> **Not needed for Phase 1–4. Add in Phase 6 before handing API to your frontend friend.**

### Approach: Code-First with swaggo

- [ ] **[`swaggo/swag`](https://github.com/swaggo/swag)** — generates OpenAPI 2.0/3.0 from Go comments
- [ ] Install: `go install github.com/swaggo/swag/cmd/swag@latest`
- [ ] Run: `swag init` — scans annotations, outputs `docs/swagger.json` + `docs/swagger.yaml`
- [ ] **Chi integration:** `github.com/swaggo/http-swagger` — serves Swagger UI at `/swagger/`

### Annotation Basics

- [ ] General API info in `main.go`:
  ```go
  // @title           Todo API
  // @version         1.0
  // @description     REST API for Todo app
  // @host            localhost:8080
  // @BasePath        /api/v1
  // @securityDefinitions.apikey BearerAuth
  // @in header
  // @name Authorization
  ```
- [ ] Handler annotations:
  ```go
  // @Summary      List todos
  // @Tags         todos
  // @Produce      json
  // @Security     BearerAuth
  // @Success      200  {array}   model.Todo
  // @Failure      401  {object}  handler.ErrorResponse
  // @Router       /todos [get]
  ```
- [ ] Document request body with `@Param body body model.CreateTodoRequest true "Todo input"`
- [ ] Re-run `swag init` whenever you change annotations

### Alternative: Spec-First with ogen

- [ ] **[`ogen`](https://github.com/ogen-go/ogen)** — write `openapi.yaml` first, generate Go server code
- [ ] Better for strict API contracts (useful when frontend dev is involved from day 1)
- [ ] More setup upfront, but the spec is the source of truth

### Sharing with Your Frontend Dev

- [ ] Export `swagger.json` or `openapi.yaml` — share the file directly
- [ ] Or point them to the running `/swagger/` UI URL during dev
- [ ] React dev can use **openapi-typescript** (`npx openapi-typescript swagger.json -o api.ts`) to generate typed API client
- [ ] Document auth clearly: which endpoints need `Authorization: Bearer <token>` header

---

## Key Libraries for This Project

| Purpose                      | Library                          |
| ---------------------------- | -------------------------------- |
| HTTP router                  | `github.com/go-chi/chi/v5`       |
| SQLite driver                | `modernc.org/sqlite`             |
| Migrations                   | `github.com/pressly/goose/v3`    |
| SQL codegen (optional)       | `github.com/sqlc-dev/sqlc`       |
| CLI                          | `github.com/spf13/cobra`         |
| Test assertions              | `github.com/stretchr/testify`    |
| Hot reload                   | `github.com/air-verse/air`       |
| Env vars                     | `github.com/joho/godotenv`       |
| Password hashing             | `golang.org/x/crypto/bcrypt`     |
| Cookie sessions              | `github.com/gorilla/sessions`    |
| CSRF protection              | `github.com/gorilla/csrf`        |
| JWT                          | `github.com/golang-jwt/jwt/v5`   |
| OpenAPI/Swagger (code-first) | `github.com/swaggo/swag`         |
| Swagger UI middleware        | `github.com/swaggo/http-swagger` |

---

## Quick Reference — PHP/TS Equivalents

| Concept           | PHP / TS                    | Go                                  |
| ----------------- | --------------------------- | ----------------------------------- |
| Class             | `class Foo {}`              | `type Foo struct {}` + methods      |
| Interface         | `interface IFoo`            | `type IFoo interface {}` (implicit) |
| Exception         | `throw new Error()`         | `return fmt.Errorf("...")`          |
| Try/catch         | `try { } catch (e) {}`      | `if err != nil { return err }`      |
| Async/await       | `async/await`, Promises     | goroutines + channels               |
| Array             | `[]` / `Array<T>`           | `[]T` (slice)                       |
| Dict/Object       | `{}` / `Record<K,V>`        | `map[K]V`                           |
| Nullable          | `T \| null` / `?T`          | pointer `*T` or zero value          |
| Optional chaining | `obj?.field`                | explicit nil check                  |
| Type assertion    | `as Type` / `(Type)$x`      | `x.(Type)`                          |
| Enum              | `enum` / `const enum`       | `const` + `iota`                    |
| Generics          | `<T>`                       | `[T any]` (Go 1.18+)                |
| Package manager   | npm / composer              | `go mod`                            |
| Module exports    | `export` / `module.exports` | UPPERCASE identifier                |
