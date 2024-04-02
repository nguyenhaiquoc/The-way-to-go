package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

/*
	config https://github.com/rs/zerolog to log message in json format and to console
*/

func Add(a, b int) int {
	log.Debug().Msgf("Adding %d and %d", a, b)
	return a + b
}

func sumInts(list ...int) (sum int) {
	// https://www.educative.io/courses/the-way-to-go/challenge-variable-number-of-arguments
	for _, v := range list {
		sum += v
	}
	return sum
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Debug().Msg("Debug message")
	log.Info().Msg("Hello, World!")
	log.Info().Msg("Hello, World!")
	Add(1, 2)
}
