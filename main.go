package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

var totalInventory int
var mu sync.Mutex
var wg sync.WaitGroup

func AddToInventory(ctx context.Context, inventory int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := process(ctx, inventory); err != nil {
			log.Print("Inventory is not valid")
		}
	}()
}

// process simulates a two-second request processing time and adds the inventory to the total if it's valid.
func process(ctx context.Context, inventory int) error {
	if inventory < 0 {
		return errors.New("inventory could not be negative")
	}

	// Simulate a two-second request processing time.
	select {
	case <-time.After(2 * time.Second):
		// It takes two seconds to process the request.

		// Use a mutex to ensure the safe access to the totalInventory variable.
		mu.Lock()
		defer mu.Unlock()

		log.Print("inventory proceed successfully, ", inventory)
		totalInventory += inventory
		return nil
	case <-ctx.Done():
		return nil
	}
}

func main() {
	log.Println("Service start!")
	ctx := context.Background()

	// Create a cancelable context to allow the program to stop gracefully on an interrupt signal.
	ctx, cancelCtx := context.WithCancel(ctx)
	cancelProcess(cancelCtx)

	addInventories(ctx)

	// Wait for all inventory addition processes to finish before exiting the program.
	wg.Wait()
}

// addInventories adds 10 inventory items to the total inventory asynchronously.
func addInventories(ctx context.Context) {
	for i := 0; i < 10; i++ {
		AddToInventory(ctx, i)
	}
}

func cancelProcess(cancelCtx context.CancelFunc) {
	// Listen for the interrupt signal to gracefully stop the program.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cancelCtx()
		log.Printf("Total inventory: %v\n", totalInventory)
		os.Exit(1)
	}()
}
