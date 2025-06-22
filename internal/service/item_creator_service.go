package service

import (
	"api/internal/repository"
	"sync"
)

func CreateItems() {

	wg := sync.WaitGroup{}

	itemChannel := make(chan repository.Item)

	wg.Add(1)
	go func() {
		TranferItems(itemChannel)
		wg.Done()
	}()

	FillItemChannel(itemChannel)

	wg.Wait()
	close(itemChannel)
}

func FillItemChannel(itemChan chan<- repository.Item) {

	repository.ReadNewUsers(itemChan)

	repository.ReadNewProducts(itemChan)

	repository.ReadNewCarts(itemChan)

	repository.ReadNewCartProducts(itemChan)

	repository.ReadNewOrders(itemChan)

	repository.ReadNewOrderProducts(itemChan)
}
