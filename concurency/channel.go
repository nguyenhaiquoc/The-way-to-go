package main

import "github.com/rs/zerolog/log"

func sendValues(myIntChannel chan int) {

	for i := 0; i < 5; i++ {
		myIntChannel <- i //sending value
	}

}

func sendValuesThenClose(myIntChannel chan int) {

	for i := 0; i < 5; i++ {
		myIntChannel <- i //sending value
	}
	// close only make sense on sender side, not receiver side
	close(myIntChannel)

}
func receiveValues() {
	myIntChannel := make(chan int)

	go sendValues(myIntChannel)

	for i := 0; i < 5; i++ {
		log.Info().Msgf("Received: %d", <-myIntChannel) //receiving value
	}
}

func receiveValuesDeadLockOnReceive() {
	myIntChannel := make(chan int)

	go sendValues(myIntChannel)

	for i := 0; i < 6; i++ {
		// Deadlock because we are trying to receive 6 values while only 5 values are sent
		// both receive and send are blocking operations
		log.Info().Msgf("Received: %d", <-myIntChannel) //receiving value
	}
}

func receiveValuesCloseChannel() {
	myIntChannel := make(chan int)

	go sendValuesThenClose(myIntChannel)

	for i := 0; i < 10; i++ {
		// alwyas return zero value when reading from closed channel
		value, open := <-myIntChannel
		if !open {
			break
		}
		log.Info().Msgf("Received: %d", value) //receiving value
	}
}

func receiveUsingRange() {
	myIntChannel := make(chan int)

	go sendValuesThenClose(myIntChannel)
	// we might want to close channel here in main goroutine

	for value := range myIntChannel {
		log.Info().Msgf("Received: %d", value) //receiving value
	}
}
