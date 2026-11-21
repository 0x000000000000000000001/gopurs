package example_http

import (
	"context"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"
)

func startTestServer(t *testing.T, handler func(Request, func(Response))) (string, context.CancelFunc, <-chan error) {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { done <- serveListener(ctx, listener, handler) }()
	t.Cleanup(func() { cancel(); _ = listener.Close() })
	return "http://" + listener.Addr().String(), cancel, done
}

func awaitStop(t *testing.T, done <-chan error) {
	t.Helper()
	select {
	case err := <-done:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(7 * time.Second):
		t.Fatal("server did not stop")
	}
}

func TestRequestTransportAndBodyLimit(t *testing.T) {
	url, cancel, done := startTestServer(t, func(request Request, reply func(Response)) {
		if request.Method != "POST" || request.Path != "/echo" || request.Body != "hello" {
			t.Error("unexpected request", request)
		}
		reply(Response{Status: 201, Body: `{"ok":true}`})
		reply(Response{Status: 500, Body: "duplicate"})
	})
	client := &http.Client{Timeout: 3 * time.Second}
	response, err := client.Post(url+"/echo", "text/plain", strings.NewReader("hello"))
	if err != nil {
		t.Fatal(err)
	}
	body, err := io.ReadAll(response.Body)
	_ = response.Body.Close()
	if err != nil || response.StatusCode != 201 || string(body) != `{"ok":true}` {
		t.Fatal(response.StatusCode, string(body), err)
	}
	response, err = client.Post(url+"/echo", "text/plain", strings.NewReader(strings.Repeat("x", (64<<10)+1)))
	if err != nil {
		t.Fatal(err)
	}
	_ = response.Body.Close()
	if response.StatusCode != http.StatusRequestEntityTooLarge {
		t.Fatal("body limit", response.StatusCode)
	}
	cancel()
	awaitStop(t, done)
}

func TestShutdownDrainsApplicationAfterClientDisconnect(t *testing.T) {
	replyReady := make(chan func(Response), 1)
	url, stop, done := startTestServer(t, func(_ Request, reply func(Response)) { replyReady <- reply })
	ctx, cancelRequest := context.WithCancel(context.Background())
	defer cancelRequest()
	request, err := http.NewRequestWithContext(ctx, "GET", url+"/pending", nil)
	if err != nil {
		t.Fatal(err)
	}
	clientDone := make(chan struct{})
	go func() {
		defer close(clientDone)
		response, _ := http.DefaultClient.Do(request)
		if response != nil {
			_ = response.Body.Close()
		}
	}()
	var reply func(Response)
	select {
	case reply = <-replyReady:
	case <-time.After(3 * time.Second):
		t.Fatal("request was not dispatched")
	}
	cancelRequest()
	<-clientDone
	stop()
	select {
	case <-done:
		t.Fatal("server returned while the application still owned resources")
	case <-time.After(25 * time.Millisecond):
	}
	reply(Response{Status: 200, Body: "{}"})
	awaitStop(t, done)
}
