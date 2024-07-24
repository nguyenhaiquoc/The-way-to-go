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
		close(channel1)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			channel2 <- "I'll print every 1s"
			time.Sleep(time.Second * 1)
		}
		close(channel2)
	}()

	for i := 0; i < 50; i++ {
		// log info i vlue to the consone
		log.Info().Msgf("i: %d", i)
		select {
		case message1, ok := <-channel1:
			if !ok {
				log.Info().Msg("Channel 1 is closed")
				// Nil channel will block the select statement forever, so this case will never be selected again
				channel1 = nil
				continue
			}
			log.Info().Msg(message1)
		case message2, ok := <-channel2:
			if !ok {
				log.Info().Msg("Channel 2 is closed")
				channel2 = nil
				continue
			}
			log.Info().Msg(message2)
		default:
			// sleep for 200ms
			time.Sleep(time.Millisecond * 200)
			// stop the select statement  if both channel1 and channel2 are closed
			if channel1 == nil && channel2 == nil {
				log.Info().Msg("Both channel1 and channel2 are closed")
				return
			}
		}

	}
}
