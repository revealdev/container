package heap_test

import (
	"github.com/revealdev/container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapify(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "already a heap",
			input:  []int{1, 3, 5, 7, 9},
			expect: []int{1, 3, 5, 7, 9},
		},
		{
			name:   "unsorted list",
			input:  []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expect: []int{1, 1, 2, 3, 3, 9, 4, 6, 5, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap.Heapify(tt.input, func(i, j int) bool { return tt.input[i] < tt.input[j] })
			for i := len(tt.input)/2 - 1; i >= 0; i-- {
				left := 2*i + 1
				right := 2*i + 2
				if left < len(tt.input) {
					assert.LessOrEqual(t, tt.input[i], tt.input[left], "Heap property violated at index %d", i)
				}
				if right < len(tt.input) {
					assert.LessOrEqual(t, tt.input[i], tt.input[right], "Heap property violated at index %d", i)
				}
			}
		})
	}
}

func TestShiftDown(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "valid shift down",
			input:  []int{10, 3, 9, 4, 8, 2},
			expect: []int{3, 4, 9, 10, 8, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := len(tt.input) - 1; i >= 0; i-- {
				heap.ShiftDown(tt.input, func(i, j int) bool { return tt.input[i] < tt.input[j] }, i)

				left := 2*i + 1
				right := 2*i + 2
				if left < len(tt.input) {
					assert.LessOrEqual(t, tt.input[i], tt.input[left], "Heap property violated at index %d", i)
				}
				if right < len(tt.input) {
					assert.LessOrEqual(t, tt.input[i], tt.input[right], "Heap property violated at index %d", i)
				}
			}
		})
	}
}

func TestPush(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		val    int
		expect []int
	}{
		{
			name:   "push into heap",
			input:  []int{1, 3, 5, 7, 9},
			val:    4,
			expect: []int{1, 3, 4, 7, 9, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap.Heapify(tt.input, func(i, j int) bool { return tt.input[i] < tt.input[j] })
			heap.Push(&tt.input, func(i, j int) bool { return tt.input[i] < tt.input[j] }, tt.val)
			for i := len(tt.input)/2 - 1; i >= 0; i-- {
				left := 2*i + 1
				right := 2*i + 2
				if left < len(tt.input) {
					assert.LessOrEqual(t, tt.input[i], tt.input[left], "Heap property violated at index %d", i)
				}
				if right < len(tt.input) {
					assert.LessOrEqual(t, tt.input[i], tt.input[right], "Heap property violated at index %d", i)
				}
			}
		})
	}
}

func TestShiftUp(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		index  int
		expect []int
	}{
		{
			name:   "shift up element",
			input:  []int{1, 3, 5, 7, 9, 4},
			index:  5,
			expect: []int{1, 3, 4, 7, 9, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap.ShiftUp(tt.input, func(i, j int) bool { return tt.input[i] < tt.input[j] }, tt.index)
			for i := len(tt.input)/2 - 1; i >= 0; i-- {
				left := 2*i + 1
				right := 2*i + 2
				if left < len(tt.input) {
					assert.LessOrEqual(t, tt.input[i], tt.input[left], "Heap property violated at index %d", i)
				}
				if right < len(tt.input) {
					assert.LessOrEqual(t, tt.input[i], tt.input[right], "Heap property violated at index %d", i)
				}
			}
		})
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect int
		remain []int
	}{
		{
			name:   "pop from heap",
			input:  []int{1, 3, 5, 7, 9},
			expect: 1,
			remain: []int{3, 5, 7, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap.Heapify(tt.input, func(i, j int) bool { return tt.input[i] < tt.input[j] })
			val, ok := heap.Pop(&tt.input, func(i, j int) bool { return tt.input[i] < tt.input[j] })
			assert.True(t, ok, "Pop operation failed")
			assert.Equal(t, tt.expect, val, "Popped value is incorrect")
			assert.Equal(t, tt.remain, tt.input, "Remaining heap is incorrect")
		})
	}
}

func TestPopPush(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		val    int
		expect int
		remain []int
	}{
		{
			name:   "pop and push in heap",
			input:  []int{1, 3, 5, 7, 9},
			val:    4,
			expect: 1,
			remain: []int{3, 4, 5, 7, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap.Heapify(tt.input, func(i, j int) bool { return tt.input[i] < tt.input[j] })
			val, ok := heap.PopPush(&tt.input, func(i, j int) bool { return tt.input[i] < tt.input[j] }, tt.val)
			assert.True(t, ok, "PopPush operation failed")
			assert.Equal(t, tt.expect, val, "Popped value is incorrect")
			assert.Equal(t, tt.remain, tt.input, "Remaining heap is incorrect")
		})
	}
}
