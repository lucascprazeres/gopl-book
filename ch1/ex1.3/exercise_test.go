package main

import "testing"

func BenchmarkLoopEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopEcho()
	}
}

func BenchmarkJoinEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoinEcho()
	}
}
