# Task list API: PureScript → Go → SQLite

A task-list JSON API with creation, listing, completion and deletion of tasks.
Routing, validation and database queries are written in PureScript.
`Database.SQLite3` provides the native database port; Argonaut handles
JSON. `Example.HTTP` is a small, ordinary Go `net/http` transport that passes typed
request/response records and callbacks through gopurs's FFI.

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

Build the gopurs JavaScript compiler as described in its bootstrap guide first.
The example needs Node, Go, Spago and the matching typed PureScript frontend.
`GOPURS_PURS` can select the frontend explicitly; the build script otherwise uses
gopurs's normal local discovery. `GOPURS_COMPILER_ROOT` and `GOPURS_LIBRARY_ROOT`
override the default sibling paths.

From this example directory:

```sh
node tools/build.mjs
TASK_LIST_DATABASE=./tasks.db TASK_LIST_ADDRESS=127.0.0.1:8080 ./.build/task-list-api
```

The build script freezes the library sources into a private temporary workspace,
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

## Verify

```sh
node tools/build.mjs --race
node --test test/http.test.mjs
node --test test/cancellation.test.mjs
go test -race -count=1 ./src/Example
```

The HTTP test runs the actual compiled executable against a temporary database
and an ephemeral loopback port. It checks CRUD, JSON validation, parameterized
SQL, concurrent insertion IDs, persistence across restart, SIGTERM shutdown and
nonzero exit status when the database or listening address cannot be opened.
The Go transport tests also check the body limit, duplicate callbacks and draining
an in-flight application operation after the client disconnects.
The cancellation test compiles a separate PureScript fixture and cancels its
server fiber while a handler is waiting at a test-controlled gate. With connected
and disconnected clients, the handler finishes before the outer resource is
released and `killFiber` returns. Its generated executable runs under `-race`.

Use `node tools/build.mjs --dev --race` for the development profile. Set
`GOPURS_JS=0` to exercise a freshly rebuilt native compiler host.

The transport owns sockets and signals; a PureScript bracket owns its native
server handle, and the outer `withDatabase` bracket owns SQLite. Canceling the
server's Aff stops waiting, then its masked release explicitly stops and joins
the server before the database can close. Every dispatched request must invoke
its reply callback, including on errors. The example's `runAff_` handler does so
on both outcomes. Long-running application work can delay shutdown beyond the
HTTP grace period; the database is kept open until that work releases its callback.

No additional HTTPurple port is needed for this example. The existing Node HTTP
port has incomplete server-lifecycle operations; this narrow local transport
keeps the example's lifecycle explicit without expanding that port's scope.
