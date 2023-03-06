package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConcurrentAddToInventory(t *testing.T) {
	//This test ensures that the AddToInventory function can add inventory items concurrently and calculate the total inventory correctly.
	// Create a context with a timeout to ensure that the test does not run indefinitely.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the number of inventory items to add in each goroutine.
	inventoryToAdd := 10

	// Add the inventory items concurrently.
	for i := 0; i < inventoryToAdd; i++ {
		AddToInventory(ctx, i)
	}

	// Wait for all goroutines to finish.
	wg.Wait()

	// Check that the total inventory is correct.
	expectedTotal := inventoryToAdd * (inventoryToAdd - 1) / 2
	if totalInventory != expectedTotal {
		t.Errorf("Expected total inventory %d, but got %d", expectedTotal, totalInventory)
	}
}

func TestCancelOnCtrlC(t *testing.T) {
	totalInventory = 0
	// This test verifies whether an interrupt signal (e.g. ctrl+c) can stop a goroutine that adds inventories and prints out the total inventory.
	var wg sync.WaitGroup
	wg.Add(1)

	ctx, cancelFunc := context.WithCancel(context.Background())

	// Create a channel to listen for interrupts (e.g. ctrl+c)
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Create a goroutine that waits for an interrupt and cancels the context
	go func() {
		<-c // Wait for a signal to be received
		cancelFunc()
	}()

	// Create another goroutine that adds inventories while the context is not cancelled
	go func() {
		defer wg.Done()     // Mark the goroutine as done when it exits
		addInventories(ctx) // Call a function to add inventories until the context is cancelled
	}()

	// Wait for a second before sending an interrupt to the channel
	time.Sleep(1 * time.Second)
	c <- syscall.SIGTERM // Send an interrupt to the channel

	wg.Wait()
	assert.Equal(t, 0, totalInventory)
}
