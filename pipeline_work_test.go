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

func BenchmarkRoutinesBC(b *testing.B) {
	for n := 0; n < b.N; n++ {
		go C(B(n))
	}
}

func BenchmarkRoutinesABC(b *testing.B) {
	for n := 0; n < b.N; n++ {
		go C(B(A(n)))
	}
}

func BenchmarkChanneledRoutineABC(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// todo
	}
}

func BenchmarkChanneledRoutineBC(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// todo
	}
}
