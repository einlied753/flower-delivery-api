package service

import (
	"api/internal/repository"
	"context"
	"fmt"
)

func TranferItems(ctx context.Context, itemChannel <-chan repository.Item) {
	for item := range itemChannel {
		select {
		case <-ctx.Done():
			fmt.Println("Context done in TranferItems()")
			return
		default:
			repository.AddItem(item)
		}
	}
}
