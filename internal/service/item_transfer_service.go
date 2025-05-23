package service

import (
	"api/internal/repository"
)

func TranferItems(itemChannel <-chan repository.Item) {
	for item := range itemChannel {
		repository.AddItem(item)
	}
}
