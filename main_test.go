package main

import (
	"context"
	"testing"
	"time"
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
