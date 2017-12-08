package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	// Mock request
	request := httptest.NewRequest("GET", "/", nil)
	// Mock response
	w := httptest.NewRecorder()
	// This is the handler() from main.go
	handler(w, request)
	// Read the result
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	// Run tests
	if resp.StatusCode != 200 {
		t.Errorf("Received status code: %d", resp.StatusCode)
	}
	if string(body) != "Hello, KubeCon!" {
		t.Errorf("Received body: %s", string(body))
	}
}
