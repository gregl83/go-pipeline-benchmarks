package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"time"
)

var (
	ResponseTime = 50 * time.Millisecond
)

// StartServer launches a new server for testing
func StartServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(ResponseTime)
	}))
}