# Task list API: PureScript → Go → SQLite

A JSON API for creating, listing, completing and deleting tasks. Use this
example to follow a request from PureScript routing through Go to SQLite.
Routing, validation and database queries are written in PureScript.
`Database.SQLite3` provides the native database port; Argonaut handles
JSON. `Example.HTTP` is a small, ordinary Go `net/http` transport that passes typed
request/response records and callbacks through
[gopurs's FFI](../../docs/ffi-idiomatic.md).

[Build and run](#build-and-run) the server, then [try the API](#try-the-api)
from another terminal. [Read the code](#read-the-code) to see where each part
lives, or [run the tests](#verify) against the compiled application.

## Build and run

Use the sibling development checkout layout:

```text
workspace/
  gopurs/
    examples/task-list-api/
  gopurs-aff/
  gopurs-strings/
  ...other Go library ports...
  gopurs-sqlite3/
```

Build the gopurs JavaScript compiler with the
[bootstrap guide](../../docs/bootstrap.md) first.
The example needs Node, Go, Spago and the matching typed PureScript frontend.
`GOPURS_PURS` can select the frontend explicitly; the build script otherwise uses
gopurs's normal local discovery. `GOPURS_COMPILER_ROOT` and `GOPURS_LIBRARY_ROOT`
override the default sibling paths.

From this example directory:

```sh
node tools/build.mjs
TASK_LIST_DATABASE=./tasks.db TASK_LIST_ADDRESS=127.0.0.1:8080 ./.build/task-list-api
```

The build script copies the library sources into a private temporary workspace,
builds the PureScript frontend, prepares the pinned Go dependencies, generates Go
and builds the executable. Logs and a build description are kept in `.build/`.
It does not clean a library checkout's output. Set `GOPURS_OFFLINE=1` for an already
cached Spago graph; the Go preparation step still needs its pinned module cache
or access to the Go module proxy. `GOPURS_TEST_KEEP_WORKSPACE=1` retains the build
workspace for inspection.

The defaults are `127.0.0.1:8080` and `tasks.db`. Port `0` selects an available
port, printed as `LISTENING http://127.0.0.1:...`. Ctrl-C or SIGTERM stops accepting
requests, drains application callbacks and closes SQLite before printing `STOPPED`.

## Try the API

```sh
curl -i http://127.0.0.1:8080/health
curl -i -H 'Content-Type: application/json' \
  -d '{"title":"Learn gopurs"}' http://127.0.0.1:8080/tasks
curl -i http://127.0.0.1:8080/tasks
curl -i http://127.0.0.1:8080/tasks/1
curl -i -X PATCH -H 'Content-Type: application/json' \
  -d '{"completed":true}' http://127.0.0.1:8080/tasks/1
curl -i -X DELETE http://127.0.0.1:8080/tasks/1
```

Use the ID returned by POST in later requests. IDs are decimal strings so JSON
does not lose the precision of SQLite's 64-bit rowids. Titles are trimmed and
must contain 1–200 UTF-16 code units. SQL uses bound parameters.

| Request | Success | Invalid or missing input |
| --- | --- | --- |
| `GET /health` | 200, `{"status":"ok"}` | — |
| `GET /tasks` | 200, array of tasks | — |
| `POST /tasks` with `{ "title": "..." }` | 201, created task | 400 |
| `GET /tasks/:id` | 200, task | 404 |
| `PATCH /tasks/:id` with `{ "completed": true }` | 200, updated task | 400 or 404 |
| `DELETE /tasks/:id` | 204 | 404 |

Unknown routes return 404. Unsupported methods on an individual task return 405.
Requests larger than 64 KiB return 413 before application dispatch. Application
errors are logged and return a generic 500 response. Restarting with the same
database path preserves data.

## Read the code

Start with `Main.purs`: it opens the database, starts the server and dispatches
each request. The HTTP files supply the transport used by that PureScript code.

| File or library | Responsibility |
| --- | --- |
| [Main.purs](src/Main.purs) | Routes, input validation, JSON responses, parameterized SQL and database lifetime. |
| [HTTP.purs](src/Example/HTTP.purs) | Request and response types, typed callbacks and the server's acquisition/release bracket. |
| [HTTP.go](src/Example/HTTP.go) | Sockets, request body limits, signals and waiting for handlers during shutdown. |
| [gopurs-sqlite3](https://github.com/0x000000000000000000001/gopurs-sqlite3) | SQLite operations, transactions and database resource management; a sibling checkout in the build layout above. |

## Verify

```sh
node tools/build.mjs --race
node --test test/http.test.mjs
node --test test/cancellation.test.mjs
go test -race -count=1 ./src/Example
```

The HTTP test runs the actual compiled executable against a temporary database
and an available local port. It checks task creation, reading, updating and
deletion, JSON validation, parameterized SQL, concurrent insertion IDs,
persistence across restart, SIGTERM shutdown and
nonzero exit status when the database or listening address cannot be opened.
The Go transport tests also check the body limit, duplicate callbacks and draining
an in-flight application operation after the client disconnects.
The cancellation test compiles a separate PureScript fixture and cancels its
server fiber while a handler is waiting at a test-controlled gate. With connected
and disconnected clients, the handler finishes before the outer resource is
released and `killFiber` returns. Its generated executable runs under `-race`.

Use `node tools/build.mjs --dev --race` for the development profile. Set
`GOPURS_JS=0` to exercise a freshly rebuilt native compiler host.

## Resource lifetime

The transport owns sockets and signals; a PureScript bracket owns its native
server handle, and the outer `withDatabase` bracket owns SQLite. A bracket pairs
resource acquisition with release, including when the action fails or is
canceled.

Canceling the server's Aff stops waiting, then its release explicitly stops
the server and waits for it to finish before the database can close. That
release is masked: cancellation cannot interrupt it. Every dispatched request
must invoke its reply callback, including on errors. The example's `runAff_`
handler does so on both outcomes. Long-running application work can delay
shutdown beyond the HTTP grace period; the database is kept open until that
work releases its callback.

This example uses a local transport because the existing Node HTTP port has
incomplete server lifecycle operations. No additional HTTPurple port is needed.
