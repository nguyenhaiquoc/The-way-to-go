package main

import "testing"

func TestFibonacyGenerator(t *testing.T) {
	c := fibonacy(100)
	receiver(c)
}
