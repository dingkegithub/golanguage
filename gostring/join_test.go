package gostring

import (
	"strings"
	"testing"
)

const (
	TstStr = "GoString"
)

func generateStringSlice(n int) []string {
	s := make([]string, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, TstStr)
	}

	return s
}

func buildString(s []string) string {
	builder := strings.Builder{}
	for _, v := range s {
		builder.WriteString(v)
	}
	return builder.String()
}

//  go test -v -test.bench Builder -test.benchmem
func BenchmarkBuilder(b *testing.B) {
	s := generateStringSlice(1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buildString(s)
	}
}
