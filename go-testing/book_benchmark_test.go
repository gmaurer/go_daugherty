package main

import (
	"testing"
)

func BenchmarkSanitizeBookTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SanitizeBookTitle("hELlO WoRlD")
	}
}
