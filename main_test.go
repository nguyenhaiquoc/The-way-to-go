package main

import "testing"

/*
	Test Add function
*/
func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", result)
	}
}

func TestSumInts(t *testing.T) {
	// create two test case for sumInts function using test table
	testCases := []struct {
		name   string
		input  []int
		result int
	}{
		{
			name:   "Test case 1",
			input:  []int{1, 2, 3},
			result: 6,
		},
		{
			name:   "Test case 2",
			input:  []int{1, 2, 3, 4, 5},
			result: 15,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := sumInts(tc.input...)
			if result != tc.result {
				t.Errorf("sumInts(%v) = %d; want %d", tc.input, result, tc.result)
			}
		})
	}
}
