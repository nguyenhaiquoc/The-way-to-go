package main

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

func fibonacy(n int) chan int {
	c := make(chan int)
	go func() {
		for i, j := 0, 1; i < n; i, j = i+j, i {
			c <- i
			// sleep for i second
			time.Sleep(time.Second * time.Duration(i))
		}
		close(c)
	}()
	return c
}

func receiver(c chan int) {
	for i := range c {
		log.Info().Msg(fmt.Sprintf("%d", i))
	}
}
