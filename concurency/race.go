package main

import (
	"github.com/rs/zerolog/log"
)

func generateDataRace() {
	i := 0
	go func() {
		i++
	}()
	i = i + 1                    // Data Race
	log.Info().Msgf("i = %d", i) // outout: i = 1 while i = 2 is expected
}

func deposit(balance *int, amount int) {
	*balance += amount //add amount to balance
}

func withdraw(balance *int, amount int) {
	*balance -= amount //subtract amount from balance
}

func generateRaceCondition() {

	balance := 100

	go deposit(&balance, 10) //depositing 10

	withdraw(&balance, 50) //withdrawing 50

	log.Info().Msgf("Balance: %d", balance) //output: Balance: 50 while Balance: 60 is expected

}
