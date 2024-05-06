package main

import (
	"context"
	"testing"

	"time"

	"github.com/rs/zerolog/log"
)

func doSomething(ctx context.Context, input chan int, done chan struct{}) {
	// Check if the context is canceled
	for {
		select {
		case <-ctx.Done():
			// Context is canceled, do cleanup or return
			log.Info().Msg("Context canceled")
			done <- struct{}{}
			return
		case i, ok := <-input:
			if !ok {
				log.Info().Msgf("Channel is closed: %d", i)
				done <- struct{}{}
				return
			}
			log.Info().Msgf("Received input: %d", i)
		}
	}
}
func TestContextCancel(t *testing.T) {
	// Create a new context with cancel function
	ctx, cancel := context.WithCancel(context.Background())

	// Use a channel to signal the completion of the goroutine
	done := make(chan struct{})
	input := make(chan int)

	// send input to the goroutine every 0.5 second
	go func() {
		for i := 0; i < 100; i++ {
			log.Info().Msgf("Sent input: %d", i)
			input <- i
			time.Sleep(300 * time.Millisecond)
			// done after 5 seconds
			/*
				if i == 5 {
					done <- struct{}{}
					return
				}
			*/
		}
		done <- struct{}{}
	}()

	// Start the goroutine
	go doSomething(ctx, input, done)

	// Cancel the context after 1 second
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	// Wait for the goroutine to complete
	<-done
	log.Info().Msg("Done")
	// close input and done channel if it is not closed
	close(done)
	close(input)
}

// test context timeout function
func TestContextTimeout(t *testing.T) {
	// Create a new context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Use a channel to signal the completion of the goroutine
	done := make(chan struct{})
	input := make(chan int)

	// send input to the goroutine every 0.5 second
	go func() {
		for i := 0; i < 100; i++ {
			log.Info().Msgf("Sent input: %d", i)
			input <- i
			time.Sleep(300 * time.Millisecond)
		}
		done <- struct{}{}
	}()

	// Start the goroutine
	go doSomething(ctx, input, done)

	// Wait for the goroutine to complete
	<-done
	log.Info().Msg("Done")
	// close input and done channel if it is not closed
	close(done)
	close(input)
}

// Todo: Nested context and cancel
