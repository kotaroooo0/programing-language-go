package main

import (
	"math/rand"
	"testing"
)

var s = randStringSlice(10, 10000)

func BenchmarkEchoFor(b *testing.B) {
	echoFor(s)
}

// benchmark
// 0.123 ns/op	       0 B/op	       0 allocs/op

func BenchmarkEchoJoin(b *testing.B) {
	echoJoin(s)
}

// benchmark
// 0.00203 ns/op	       0 B/op	       0 allocs/op

var rs1Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(b)
}

func randStringSlice(n, s int) []string {
	randSlice := make([]string, s)
	for i := 0; i < s; i++ {
		randSlice = append(randSlice, randString(n))
	}
	return randSlice
}
