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

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Debug().Msg("Debug message")
	log.Info().Msg("Hello, World!")
	log.Info().Msg("Hello, World!")
	Add(1, 2)
}
