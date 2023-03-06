package main

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

var totalInventory int
var mu sync.Mutex
var wg sync.WaitGroup

func AddToInventory(_ context.Context, inventory int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := process(inventory); err != nil {
			log.Print("Inventory is not valid")
		}
	}()
}

func process(inventory int) error {
	if inventory < 0 {
		return errors.New("inventory could not be negative")

	}
	// It takes two seconds to process the request.
	time.Sleep(time.Second * 2)

	mu.Lock()
	defer mu.Unlock()

	log.Print("inventory proceed successfully, ", inventory)
	totalInventory += inventory
	return nil
}

func main() {
	ctx := context.Background()
	addInventories(ctx)
	wg.Wait()
}

func addInventories(ctx context.Context) {
	for i := 0; i < 10; i++ {
		AddToInventory(ctx, i)
	}
}
