package main

import (
	"testing"
)

func BenchmarkSynchronousABC(b *testing.B) {
	for n := 0; n < b.N; n++ {
		C(B(A(n)))
	}
}

func BenchmarkSynchronousBC(b *testing.B) {
	for n := 0; n < b.N; n++ {
		C(B(n))
	}
}

func BenchmarkRoutinesABC(b *testing.B) {
	var n int
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			n = n + 1
			go C(B(A(n)))
		}
	})
}

func BenchmarkRoutinesBC(b *testing.B) {
	var n int
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			n = n + 1
			go C(B(n))
		}
	})
}

func BenchmarkChanneledRoutineABC(b *testing.B) {
	in := make(chan Payload)
	out := make(chan Payload)
	defer close(in)
	defer close(out)

	aOut := make(chan Payload)
	defer close(aOut)

	go func(in chan Payload, out chan Payload) {
		for payload := range in {
			payload.Data.Iteration = A(payload.Data.Iteration)
			out <- payload
		}
	}(in, aOut)

	bOut := make(chan Payload)
	defer close(bOut)

	go func(in chan Payload, out chan Payload) {
		for payload := range in {
			payload.Data.Hash = B(payload.Data.Iteration)
			out <- payload
		}
	}(aOut, bOut)

	go func(in chan Payload, out chan Payload) {
		for payload := range in {
			C(payload.Data.Hash)
			out <- payload
		}
	}(bOut, out)

	var n int
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			n = n + 1
			payload := Payload{
				Data: Data{
					Iteration: n,
				},
			}
			in <- payload
			// todo make sure this is actually running in parallel
			<-out
		}
	})
}

func BenchmarkChanneledRoutineBC(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// todo
	}
}