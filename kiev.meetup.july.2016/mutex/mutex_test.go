package main

import "testing"

func BenchmarkMux(b *testing.B) {
	c := NewCounterMux("")
	for i := 0; i < b.N; i++ {
		c.Add(42)
	}
}

func BenchmarkAtomic(b *testing.B) {
	c := NewCounter("")
	for i := 0; i < b.N; i++ {
		c.Add(42)
	}
}
