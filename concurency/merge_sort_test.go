package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "Sorted input",
			input:  []int{1, 2},
			output: []int{1, 2},
		},
		{
			name:   "Reverse sorted input",
			input:  []int{4, 3, 2, 1},
			output: []int{1, 2, 3, 4},
		},
		{
			name:   "Random input",
			input:  []int{3, 1, 4, 2, 5},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "Random input",
			input:  []int{3, 1, 4, 2, 6, 5},
			output: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "Empty input",
			input:  []int{},
			output: []int{},
		},
		{
			name:   "One element input",
			input:  []int{10},
			output: []int{10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeSort(tt.input)
			if !reflect.DeepEqual(result, tt.output) {
				t.Errorf("Expected %v, but got %v", tt.output, result)
			}
		})
	}
}

func TestMergeSortConcurrent(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "Sorted input",
			input:  []int{1, 2},
			output: []int{1, 2},
		},
		{
			name:   "Reverse sorted input",
			input:  []int{4, 3, 2, 1},
			output: []int{1, 2, 3, 4},
		},
		{
			name:   "Random input",
			input:  []int{3, 1, 4, 2, 5},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "Random input",
			input:  []int{3, 1, 4, 2, 6, 5},
			output: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "Empty input",
			input:  []int{},
			output: []int{},
		},
		{
			name:   "One element input",
			input:  []int{10},
			output: []int{10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeSortConcurent(tt.input)
			if !reflect.DeepEqual(result, tt.output) {
				t.Errorf("Expected %v, but got %v", tt.output, result)
			}
		})
	}
}
