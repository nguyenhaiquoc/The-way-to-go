package coffee

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.With().Caller().Logger()
	// log current log level
	log.Info().Msgf("Current log level: %v", zerolog.GlobalLevel())
	Coffees = CoffeeList{
		List: []CoffeeDetails{
			{"Latte", 2.5},
			{"Flat White", 2},
			{"Cappucinno", 2.25},
		},
	}
}

func TestIsCoffeeAvailable(t *testing.T) {
	type testCase struct {
		coffeeType string
		want       bool
	}

	cases := []testCase{
		{"lat", false},
		{"Latte", true},
		{"", false},
		{"cappacunio", false},
	}
	// zeloglog to log confee list
	log.Debug().Msgf("Coffee list: %v", Coffees)

	for _, tc := range cases {
		got := IsCoffeeAvailable(tc.coffeeType)
		if tc.want != got {
			t.Errorf("Expected '%v', but got '%v'", tc.want, got)
		}
	}
}
