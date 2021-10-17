package main

import (
	"os"
	"testing"
)

func Benchmark_findFilesWithWalkDir(b *testing.B) {
	SearchPath := os.ExpandEnv("${GOPATH}")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findFilesWithWalkDir(SearchPath)
	}
}

func Benchmark_findFilesWithWalk(b *testing.B) {
	SearchPath := os.ExpandEnv("${GOPATH}")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findFilesWithWalk(SearchPath)
	}
}
