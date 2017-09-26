package main

import (
	"fmt"
	"strconv"
	"crypto/sha256"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Payload struct {
	Iteration int
}

var (
	server = CreateServer()
)

func A(iteration int) int {
	reader, writer := io.Pipe()
	go func(iteration int) {
		json.NewEncoder(writer).Encode(Payload{
			Iteration: iteration,
		})
		writer.Close()
	}(iteration)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", server.URL, reader)
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	resIteration, _ := strconv.Atoi(string(body))
	return resIteration
}

func B(iteration int) string {
	h := sha256.New()
	h.Write([]byte(string(iteration)))
	return string(h.Sum(nil))
}

func C(result string) {
	fmt.Sprintf("result: %s\n\r", result)
}