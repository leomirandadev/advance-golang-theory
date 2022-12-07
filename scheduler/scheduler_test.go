package scheduler

import "testing"

func BenchmarkTask1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Task1()
	}
}

func BenchmarkExecCalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		execCalc()
	}
}

func BenchmarkTask2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Task2()
	}
}
