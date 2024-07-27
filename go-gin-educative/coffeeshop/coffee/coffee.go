package coffee

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

//var coffees = map[string]float32{"Latte":2.5, "Cappuccino": 2.75, "Flat White": 2.25}

type CoffeeDetails struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type CoffeeList struct {
	List []CoffeeDetails `json:"list"`
}

var Coffees CoffeeList

func GetCoffees() (*CoffeeList, error) {
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: %w", err)
		return nil, err
	}
	// log raw data viper read from config file
	log.Debug().Msgf("Raw data: %v", viper.AllSettings())
	// Log all keys viper read from config file in json format
	log.Debug().Msgf("All keys: %v", viper.AllKeys())

	// unmarshal data into coffees
	err = viper.Unmarshal(&Coffees)
	if err != nil {
		return nil, err
	}
	return &Coffees, nil
}

func IsCoffeeAvailable(coffeetype string) bool {
	for _, element := range Coffees.List {
		if element.Name == coffeetype {
			result := fmt.Sprintf("%s for $%v", element.Name, element.Price)
			log.Debug().Msgf("Coffee found: %v", result)
			return true
		}
	}
	return false

}
