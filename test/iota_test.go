package test

import (
	"fmt"
	"runtime"
	"strconv"

	"testing"
)

func init() {
	runtime.GOMAXPROCS(4)
}

func TestSprintf(t *testing.T) {
	num := 20
	t.Parallel()
	if testing.Short() {
		fmt.Println(num)
	} else {
		fmt.Println(num)
		fmt.Println(num)
	}
}

func BenchmarkSprintf(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", num)
	}
}

func BenchmarkFormat(b *testing.B) {
	num := int64(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(num, 10)
	}
}
