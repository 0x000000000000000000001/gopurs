package purescript

import "gopurs/output/gopurs_runtime"

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Request struct {
	Method string
	Path   string
	Body   string
}

type Response struct {
	Status int64
	Body   string
}

type Server struct {
	cancel context.CancelFunc
	done   chan struct{}
	err    error
}

// The application owns routing, validation, JSON and persistence. This file
// only transports closed records and callbacks across the ordinary Go FFI.
func Example_HTTP_StartImpl(ctx context.Context, address string, handler func(Request, func(Response))) (*Server, func(), error) {
	if err := ctx.Err(); err != nil {
		return nil, nil, err
	}
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, nil, err
	}
	// Acquisition's bridge context ends as soon as the handle is delivered.
	// The server owns a separate lifetime; StopImpl joins it during release.
	lifetime, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	server := &Server{cancel: stop, done: make(chan struct{})}
	go func() {
		defer close(server.done)
		defer stop()
		defer listener.Close()
		server.err = serveListener(lifetime, listener, handler)
	}()
	return server, func() { _ = Example_HTTP_StopImpl(server) }, nil
}

func Example_HTTP_WaitImpl(ctx context.Context, server *Server) error {
	select {
	case <-server.done:
		return server.err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func Example_HTTP_StopImpl(server *Server) error {
	server.cancel()
	<-server.done
	return server.err
}

func serveListener(ctx context.Context, listener net.Listener, handler func(Request, func(Response))) error {
	var requests sync.WaitGroup
	var gate sync.Mutex
	stopping := false
	server := &http.Server{
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
	server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gate.Lock()
		if stopping {
			gate.Unlock()
			http.Error(w, "server stopping", http.StatusServiceUnavailable)
			return
		}
		requests.Add(1)
		gate.Unlock()
		defer requests.Done()
		defer r.Body.Close()
		body, readErr := io.ReadAll(http.MaxBytesReader(w, r.Body, 64<<10))
		if readErr != nil {
			var sizeErr *http.MaxBytesError
			status := http.StatusBadRequest
			if errors.As(readErr, &sizeErr) {
				status = http.StatusRequestEntityTooLarge
			}
			http.Error(w, http.StatusText(status), status)
			return
		}

		responses := make(chan Response, 1)
		var once sync.Once
		handler(Request{Method: r.Method, Path: r.URL.Path, Body: string(body)}, func(response Response) {
			once.Do(func() { responses <- response })
		})
		// Wait even after client disconnection: the PureScript Aff may still be
		// using SQLite. Shutdown must finish that work before closing the DB.
		response := <-responses
		status := int(response.Status)
		if status < 200 || status > 599 {
			status = http.StatusInternalServerError
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		w.WriteHeader(status)
		_, _ = io.WriteString(w, response.Body)
	})

	served := make(chan error, 1)
	go func() { served <- server.Serve(listener) }()
	fmt.Printf("LISTENING http://%s\n", listener.Addr())
	var err error
	select {
	case err = <-served:
	case <-ctx.Done():
	}
	gate.Lock()
	stopping = true
	gate.Unlock()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	shutdownErr := server.Shutdown(shutdownCtx)
	if shutdownErr != nil {
		_ = server.Close()
	}
	requests.Wait()
	if shutdownErr != nil {
		return shutdownErr
	}
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Example_HTTP_StartImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.EffectFnAffAcquire(func(ctx gopurs_runtime.FFIContext) (gopurs_runtime.Value, func(), error) {
		go_arg0 := gopurs_runtime.Unbox[string](arg0)
		go_arg1 := func(p0_0 Request, p0_1 func(Response)) {
			gopurs_runtime.Apply2(arg1, func() gopurs_runtime.Value {
				_source := p0_0
				_ = _source
				return gopurs_runtime.RecordDict3("method", "path", "body", gopurs_runtime.Box(_source.Method), gopurs_runtime.Box(_source.Path), gopurs_runtime.Box(_source.Body))
			}(), gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				p0_1(func() Response {
					_source := arg
					_ = _source
					return Response{Status: gopurs_runtime.Unbox[int64](gopurs_runtime.RecordGet(_source, "status")), Body: gopurs_runtime.Unbox[string](gopurs_runtime.RecordGet(_source, "body"))}
				}())
				return gopurs_runtime.Value{}
			}))
		}
		go_res, go_cleanup, go_err := Example_HTTP_StartImpl(ctx, go_arg0, go_arg1)
		if go_err != nil {
			return gopurs_runtime.Value{}, go_cleanup, go_err
		}
		return gopurs_runtime.Box(go_res), go_cleanup, nil
	})
})
var _Gopurs_Example_HTTP_StopImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.EffectFnAff(func(ctx gopurs_runtime.FFIContext) (gopurs_runtime.Value, error) {
		var go_arg0 *Server
		if arg0.Type == gopurs_runtime.TypeAny {
			go_arg0 = arg0.AnyVal().(*Server)
		} else {
			go_arg0 = *(**Server)(arg0.UnsafePtr)
		}
		go_err := Example_HTTP_StopImpl(go_arg0)
		return gopurs_runtime.Value{}, go_err
	})
})
var _Gopurs_Example_HTTP_WaitImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.EffectFnAff(func(ctx gopurs_runtime.FFIContext) (gopurs_runtime.Value, error) {
		var go_arg0 *Server
		if arg0.Type == gopurs_runtime.TypeAny {
			go_arg0 = arg0.AnyVal().(*Server)
		} else {
			go_arg0 = *(**Server)(arg0.UnsafePtr)
		}
		go_err := Example_HTTP_WaitImpl(ctx, go_arg0)
		return gopurs_runtime.Value{}, go_err
	})
})
