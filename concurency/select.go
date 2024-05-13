package main

import (
	"time"

	"github.com/rs/zerolog/log"
)

func playSelect() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			channel1 <- "I'll print every 100ms"
			time.Sleep(time.Millisecond * 100)

		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			channel2 <- "I'll print every 1s"
			time.Sleep(time.Second * 1)

		}
	}()

	for i := 0; i < 100; i++ {
		// log info i vlue to the consone
		log.Info().Msgf("i: %d", i)
		select {
		case message1 := <-channel1:
			log.Info().Msg(message1)
		case message2 := <-channel2:
			log.Info().Msg(message2)
		default: // will block the program to wait for channel signnal if default is not included
			// default is a kind of non blocking select
			log.Info().Msg("No message received")
		}
	}
}
