package service

import (
	"api/internal/repository"
	"context"
	"fmt"
	"sync"
)

func CreateItems(ctx context.Context) {

	wg := sync.WaitGroup{}
	itemChannel := make(chan repository.Item)

	wg.Add(2)

	go func() {
		defer wg.Done()
		TranferItems(itemChannel)
	}()

	go func() {
		defer wg.Done()
		FillItemChannel(ctx, itemChannel)
		close(itemChannel)
	}()

	wg.Wait()

}

func FillItemChannel(ctx context.Context, itemChan chan<- repository.Item) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context done in FillItemChannel()")
			return
		default:
			repository.ReadNewUsers(itemChan)

			repository.ReadNewProducts(itemChan)

			repository.ReadNewCarts(itemChan)

			repository.ReadNewCartProducts(itemChan)

			repository.ReadNewOrders(itemChan)

			repository.ReadNewOrderProducts(itemChan)
		}
	}
}
