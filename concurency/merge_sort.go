package main

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	left := MergeSort(data[:mid])
	right := MergeSort(data[mid:])
	return merge(left, right)
}

func MergeSortConcurent(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	left := make([]int, 0, len(data)/2)
	done := make(chan struct{})
	go func() {
		left = MergeSortConcurent(data[:mid])
		done <- struct{}{}
	}()
	right := MergeSortConcurent(data[mid:])
	<-done
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}
