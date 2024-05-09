package main

import (
	"testing"
)

// Can run race detector with go test --race
func TestGenerateDataRace(t *testing.T) {
	generateDataRace()
}

func TestGenerateRaceCondition(t *testing.T) {
	generateRaceCondition()
}
