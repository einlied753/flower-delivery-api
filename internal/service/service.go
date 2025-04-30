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
		repository.SetItems(itemChannel)
		wg.Done()
	}()

	itemSlice := FillItemSlice()

	for _, v := range itemSlice {
		itemChannel <- v
	}
	close(itemChannel)

	wg.Wait()
}

func FillItemSlice() []repository.Item {

	slice := []repository.Item{}

	newUser := user.NewUser("Annette Johnson", "ann@mail.com", "+34563456345", "34, Leroy str., 153")

	newProduct1 := product.NewProduct(1, "Red rose", 59.99, 1000, 0.5)
	newProduct2 := product.NewProduct(2, "White tulip", 39.99, 2000, 0.5)

	newCart := cart.NewCart(1, 1)

	newCartProduct1 := cart.NewCartProduct(1, 1, 1, 10)
	newCartProduct2 := cart.NewCartProduct(2, 1, 2, 15)

	newOrder := order.NewOrder(1, 299.95, 10, "ann@mail.com", "45, Apple str., 23", "+345686565665")

	newOrderProduct := order.NewOrderProduct(1, 1, 10)

	slice = append(slice, newUser, newProduct1, newProduct2, newCart, newCartProduct1, newCartProduct2, newOrder, newOrderProduct)

	return slice
}
