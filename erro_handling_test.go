package main

import (
	"testing"

	"github.com/rs/zerolog/log"
)

func TestConvertStringToint(t *testing.T) {
	s := "a123ad3234"
	x, _ := ConvertStringToint(s)
	log.Info().Msgf("The value of x is %d", x)
}
