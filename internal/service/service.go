package service

import (
	"api/internal/model/cart"
	"api/internal/model/order"
	"api/internal/model/product"
	"api/internal/model/user"
	"api/internal/repository"
	"time"
)

func Creator() {

	time.Sleep(time.Second)
	newSlice := []repository.Item{}

	newUser := user.NewUser("Annette Johnson", "ann@mail.com", "+34563456345", "34, Leroy str., 153")
	newSlice = append(newSlice, newUser)

	newProduct1 := product.NewProduct(1, "Red rose", 59.99, 1000, 0.5)
	newProduct2 := product.NewProduct(2, "White tulip", 39.99, 2000, 0.5)
	newSlice = append(newSlice, newProduct1, newProduct2)

	newCart := cart.NewCart(1, 1)
	newSlice = append(newSlice, newCart)

	newCartProduct1 := cart.NewCartProduct(1, 1, 1, 10)
	newCartProduct2 := cart.NewCartProduct(2, 1, 2, 15)
	newSlice = append(newSlice, newCartProduct1, newCartProduct2)

	newOrder := order.NewOrder(1, 299.95, 10, "ann@mail.com", "45, Apple str., 23", "+345686565665")
	newSlice = append(newSlice, newOrder)

	newOrderProduct := order.NewOrderProduct(1, 1, 10)
	newSlice = append(newSlice, newOrderProduct)

	repository.SetItems(newSlice)
}
