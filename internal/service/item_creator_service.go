package service

import (
	"api/internal/model/cart"
	"api/internal/model/order"
	"api/internal/model/product"
	"api/internal/model/user"
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

	newUser1 := user.NewUser(1, "Annette Johnson", "ann@mail.com", "+34563456345", "34, Leroy str., 153")
	itemChan <- newUser1

	newProduct1 := product.NewProduct(1, "Red rose", 59.99, 1000, 0.5)
	itemChan <- newProduct1

	newCart := cart.NewCart(1, 1)
	itemChan <- newCart

	newUser2 := user.NewUser(2, "Tom Felps", "tj@gmail.com", "+8765435678654", "25, Starter ave., 87")
	itemChan <- newUser2

	newCartProduct1 := cart.NewCartProduct(1, 1, 1, 10)
	itemChan <- newCartProduct1

	newOrder := order.NewOrder(1, 299.95, 10, "ann@mail.com", "45, Apple str., 23", "+345686565665")
	itemChan <- newOrder

	newProduct2 := product.NewProduct(2, "White tulip", 39.99, 2000, 0.5)
	itemChan <- newProduct2

	newUser3 := user.NewUser(3, "Kate Watson", "katem@gmail.com", "+764356754564", "12, Main ave., 48")
	itemChan <- newUser3

	newProduct3 := product.NewProduct(3, "Blue daisy", 49.99, 3000, 0.5)
	itemChan <- newProduct3

	newProduct4 := product.NewProduct(4, "White chrysanthemum", 79.99, 2500, 0.5)
	itemChan <- newProduct4

	newOrderProduct1 := order.NewOrderProduct(1, 1, 10)
	itemChan <- newOrderProduct1

	newCartProduct2 := cart.NewCartProduct(2, 1, 2, 15)
	itemChan <- newCartProduct2
}
