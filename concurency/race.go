package main

import "github.com/rs/zerolog/log"

func GenerateDataRace() {
	i := 0
	go func() {
		i++
	}()
	i = i + 1                    // Data Race
	log.Info().Msgf("i = %d", i) // outout: i = 1 while i = 2 is expected
}
