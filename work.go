package main

import (
	"fmt"
	"crypto/sha256"
	"time"
)

type Data struct {
	Iteration int
	Hash string
}

func A(iteration int) int {
	time.Sleep(50 * time.Millisecond)
	return iteration ^ 10
}

func B(iteration int) string {
	h := sha256.New()
	h.Write([]byte(string(iteration)))
	return string(h.Sum(nil))
}

func C(result string) string {
	return fmt.Sprintf("result: %s\n\r", result)
}