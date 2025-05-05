package repository

import (
	"api/internal/model/cart"
	"api/internal/model/order"
	"api/internal/model/product"
	"api/internal/model/user"
	"fmt"
	"sync"
)

var (
	userSlice         []*user.User
	productSlice      []*product.Product
	orderSlice        []*order.Order
	orderProductSlice []*order.OrderProduct
	cartSlice         []*cart.Cart
	cartProductSlice  []*cart.CartProduct

	userMutex         sync.RWMutex
	productMutex      sync.RWMutex
	orderMutex        sync.RWMutex
	orderProductMutex sync.RWMutex
	cartMutex         sync.RWMutex
	cartProductMutex  sync.RWMutex
)

type Item interface {
	SaveItemLog()
}

func AddItem(item Item) {

	switch v := item.(type) {
	case *user.User:
		userMutex.Lock()
		defer userMutex.Unlock()
		userSlice = append(userSlice, v)
	case *product.Product:
		productMutex.Lock()
		defer productMutex.Unlock()
		productSlice = append(productSlice, v)
	case *cart.Cart:
		cartMutex.Lock()
		defer cartMutex.Unlock()
		cartSlice = append(cartSlice, v)
	case *cart.CartProduct:
		cartProductMutex.Lock()
		defer cartProductMutex.Unlock()
		cartProductSlice = append(cartProductSlice, v)
	case *order.Order:
		orderMutex.Lock()
		defer orderMutex.Unlock()
		orderSlice = append(orderSlice, v)
	case *order.OrderProduct:
		orderProductMutex.Lock()
		defer orderProductMutex.Unlock()
		orderProductSlice = append(orderProductSlice, v)
	default:
		fmt.Println("Error in repository.SetItems: undefined type of item")
	}
}

func GetUsers() []*user.User {
	userMutex.RLock()
	defer userMutex.RUnlock()
	return userSlice
}

func GetProducts() []*product.Product {
	productMutex.RLock()
	defer productMutex.RUnlock()
	return productSlice
}

func GetCarts() []*cart.Cart {
	cartMutex.RLock()
	defer cartMutex.RUnlock()
	return cartSlice
}

func GetCartProducts() []*cart.CartProduct {
	cartMutex.RLock()
	defer cartMutex.RUnlock()
	return cartProductSlice
}

func GetOrders() []*order.Order {
	orderMutex.RLock()
	defer orderMutex.RUnlock()
	return orderSlice
}

func GetOrderProducts() []*order.OrderProduct {
	orderProductMutex.RLock()
	defer orderProductMutex.RUnlock()
	return orderProductSlice
}
