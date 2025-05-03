package service

import (
	"api/internal/repository"
	"sync"
)

func TranferItems(itemChannel <-chan repository.Item) {

	mu := sync.Mutex{}

	for item := range itemChannel {
		mu.Lock()
		repository.AddItem(item)
		mu.Unlock()
	}
}
