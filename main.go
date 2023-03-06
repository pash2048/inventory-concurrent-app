package main

import (
	"context"
	"errors"
	"log"
	"time"
)

var totalInventory int

func AddToInventory(_ context.Context, inventory int) {
	if err := process(inventory); err != nil {
		log.Print("Inventory is not valid")
	}
}

func process(inventory int) error {
	if inventory < 0 {
		return errors.New("inventory could not be negative")
	}

	// It takes two seconds to process the request.
	time.Sleep(time.Second * 2)

	log.Print("inventory proceed successfully")
	totalInventory += inventory
	return nil
}

func main() {
	ctx := context.Background()
	addInventories(ctx)
}

func addInventories(ctx context.Context) {
	for i := 0; i < 10; i++ {
		AddToInventory(ctx, i)
	}
}
