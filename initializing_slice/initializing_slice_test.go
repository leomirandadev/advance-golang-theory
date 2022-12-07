package scheduler

import (
	"math/rand"
	"testing"
)

const times = 50

func BenchmarkTask1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiplier := rand.Intn(5-1) + 1
		Task1(times, multiplier)
	}
}
func BenchmarkTask2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiplier := rand.Intn(5-1) + 1
		Task2(times, multiplier)
	}
}

func BenchmarkTask3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiplier := rand.Intn(5-1) + 1
		Task3(times, multiplier)
	}
}

func BenchmarkTask4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiplier := rand.Intn(5-1) + 1
		Task4(multiplier)
	}
}
