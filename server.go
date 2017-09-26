package main

import (
	"net/http/httptest"
	"net/http"
	"time"
	"io/ioutil"
)

var (
	ResponseTime = 50 * time.Millisecond
)

// CreateServer produces a new httptest server
func CreateServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(ResponseTime)
		body, _ := ioutil.ReadAll(r.Body)
		w.Write(body)
	}))
}