package main

import (
	"github.com/rs/zerolog/log"
)

func nestedFunction() {
	defer func() {
		log.Info().Msgf("Deferred nestedFunction called")
		// recover()
	}()
	panic("I am a panic")
}
func ConvertStringToint(s string) (int, error) {
	defer func() {
		log.Info().Msgf("Deferred function called")
		recover()
		log.Info().Msgf("Recovered from panic")
	}()
	var i int
	log.Info().Msgf("start running ConvertStringToint")
	nestedFunction()
	log.Info().Msgf("continue running ConvertStringToint")
	/*
		_, err := fmt.Sscanf(s, "%d", &i)
		if err != nil {
			panic(err)
		}
	*/
	return i, nil
}
