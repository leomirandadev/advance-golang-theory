package string_append

import (
	"testing"
)

const times = 1_000

func BenchmarkAppendString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendString(times)
	}
}
func BenchmarkAppendStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendStringBuilder(times)
	}
}
