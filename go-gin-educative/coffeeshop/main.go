package main

import (
	"coffeeshop/coffee"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Coffeeshop!",
		})
	})
	r.GET("/coffee", getCoffee)
	r.Run(":8081")
}

func getCoffee(c *gin.Context) {
	coffeelist, err := coffee.GetCoffees()
	if err != nil {
		log.Error().Msgf("Error getting coffee list: %v", err)
		c.String(http.StatusInternalServerError, "Error getting coffee list")
		return
	}
	log.Debug().Msgf("Coffees: %v", coffeelist)
	c.String(http.StatusOK, " %s", coffeelist)
}
