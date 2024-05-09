package main

import "testing"

func TestSendAndReceiveValues(t *testing.T) {
	receiveValues()
}

func TestReceiveValuesDeadLockOnReceive(t *testing.T) {
	receiveValuesDeadLockOnReceive()
}

func TestReceiveValuesCloseChannel(t *testing.T) {
	receiveValuesCloseChannel()
}

func TestReceiveUsingRange(t *testing.T) {
	receiveUsingRange()
}
