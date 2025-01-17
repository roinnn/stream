package stream

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

func BenchmarkParallelByCPU(b *testing.B) {
	tests := []struct {
		name       string
		goroutines int
		action     func(int, int)
	}{
		{name: "no parallel", goroutines: 0},
		{name: "goroutines", goroutines: 2},
		{name: "goroutines", goroutines: 4},
		{name: "goroutines", goroutines: 6},
		{name: "goroutines", goroutines: 8},
		{name: "goroutines", goroutines: 10},
	}
	s := newArray(100)

	for _, tt := range tests {
		b.Run(fmt.Sprintf("%s(%d)", tt.name, tt.goroutines), func(b *testing.B) {
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = NewSlice(s).Parallel(tt.goroutines).ForEach(func(i int, v int) {
					sort.Ints(newArray(1000)) // Simulate time-consuming CPU operations
				})
			}
		})
	}
}

func BenchmarkParallelByIO(b *testing.B) {
	tests := []struct {
		name       string
		goroutines int
		action     func(int, int)
	}{
		{name: "no parallel", goroutines: 0},
		{name: "goroutines", goroutines: 2},
		{name: "goroutines", goroutines: 4},
		{name: "goroutines", goroutines: 6},
		{name: "goroutines", goroutines: 8},
		{name: "goroutines", goroutines: 10},
		{name: "goroutines", goroutines: 50},
		{name: "goroutines", goroutines: 100},
	}
	s := newArray(100)

	for _, tt := range tests {
		b.Run(fmt.Sprintf("%s(%d)", tt.name, tt.goroutines), func(b *testing.B) {
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = NewSlice(s).Parallel(tt.goroutines).ForEach(func(i int, v int) {
					time.Sleep(time.Millisecond) // Simulate time-consuming IO operations
				})
			}
		})
	}
}
