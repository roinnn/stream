package stream

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceOrderedMax(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "normal",
			input: []int{1, 2, 1, 5},
			want:  5,
		},
		{
			name:  "normal",
			input: []int{10, 2, 1, 5},
			want:  10,
		},
		{
			name:  "empty",
			input: []int{},
			want:  0,
		},
		{
			name:  "nil",
			input: nil,
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSliceByOrdered(tt.input).Max()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSliceOrderedMin(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "normal",
			input: []int{1, 2, 1, 5},
			want:  1,
		},
		{
			name:  "normal",
			input: []int{10, 2, 3, 1},
			want:  1,
		},
		{
			name:  "empty",
			input: []int{},
			want:  0,
		},
		{
			name:  "nil",
			input: nil,
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSliceByOrdered(tt.input).Min()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSliceOrderedSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "normal",
			input: []int{1, 2, 1, 5},
			want:  []int{1, 1, 2, 5},
		},
		{
			name:  "empty",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "nil",
			input: nil,
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSliceByOrdered(tt.input).Sort()
			assert.Equal(t, tt.want, got.slice)
		})
	}
}
