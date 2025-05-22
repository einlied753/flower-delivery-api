package service

import (
	"api/internal/model/cart"
	"api/internal/model/order"
	"api/internal/model/product"
	"api/internal/model/user"
	"api/internal/repository"
	"context"
	"sync"
	"time"
)

func CreateItems(ctx context.Context) {

	wg := sync.WaitGroup{}
	itemChannel := make(chan repository.Item)

	wg.Add(2)

	go func() {
		defer wg.Done()
		TranferItems(ctx, itemChannel)
	}()

	go func() {
		defer wg.Done()
		FillItemChannel(itemChannel)
		close(itemChannel)
	}()

	wg.Wait()

}

func FillItemChannel(itemChan chan<- repository.Item) {

	newUser1 := user.NewUser(1, "Annette Johnson", "ann@mail.com", "+34563456345", "34, Leroy str., 153")
	itemChan <- newUser1
	time.Sleep(100 * time.Millisecond)

	newProduct1 := product.NewProduct(1, "Red rose", 59.99, 1000, 0.5)
	itemChan <- newProduct1
	time.Sleep(100 * time.Millisecond)

	newCart := cart.NewCart(1, 1)
	itemChan <- newCart
	time.Sleep(100 * time.Millisecond)

	newUser2 := user.NewUser(2, "Tom Felps", "tj@gmail.com", "+8765435678654", "25, Starter ave., 87")
	itemChan <- newUser2
	time.Sleep(100 * time.Millisecond)

	newCartProduct1 := cart.NewCartProduct(1, 1, 1, 10)
	itemChan <- newCartProduct1
	time.Sleep(100 * time.Millisecond)

	newOrder := order.NewOrder(1, 299.95, 10, "ann@mail.com", "45, Apple str., 23", "+345686565665")
	itemChan <- newOrder
	time.Sleep(100 * time.Millisecond)

	newProduct2 := product.NewProduct(2, "White tulip", 39.99, 2000, 0.5)
	itemChan <- newProduct2
	time.Sleep(100 * time.Millisecond)

	newOrderProduct1 := order.NewOrderProduct(1, 1, 10)
	itemChan <- newOrderProduct1
	time.Sleep(100 * time.Millisecond)

	newCartProduct2 := cart.NewCartProduct(2, 1, 2, 15)
	itemChan <- newCartProduct2
	time.Sleep(100 * time.Millisecond)
}
