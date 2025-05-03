package repository

import (
	"api/internal/model/cart"
	"api/internal/model/order"
	"api/internal/model/product"
	"api/internal/model/user"
	"fmt"
)

var (
	userSlice         []*user.User
	productSlice      []*product.Product
	orderSlice        []*order.Order
	orderProductSlice []*order.OrderProduct
	cartSlice         []*cart.Cart
	cartProductSlice  []*cart.CartProduct
)

type Item interface {
	SaveItemLog()
}

func AddItem(item Item) {

	switch v := item.(type) {
	case *user.User:
		userSlice = append(userSlice, v)
	case *product.Product:
		productSlice = append(productSlice, v)
	case *cart.Cart:
		cartSlice = append(cartSlice, v)
	case *cart.CartProduct:
		cartProductSlice = append(cartProductSlice, v)
	case *order.Order:
		orderSlice = append(orderSlice, v)
	case *order.OrderProduct:
		orderProductSlice = append(orderProductSlice, v)
	default:
		fmt.Println("Error in repository.SetItems: undefined type of item")
	}
}

func GetUsers() []*user.User {
	return userSlice
}

func GetProducts() []*product.Product {
	return productSlice
}

func GetCarts() []*cart.Cart {
	return cartSlice
}

func GetCartProducts() []*cart.CartProduct {
	return cartProductSlice
}

func GetOrders() []*order.Order {
	return orderSlice
}

func GetOrderProducts() []*order.OrderProduct {
	return orderProductSlice
}
