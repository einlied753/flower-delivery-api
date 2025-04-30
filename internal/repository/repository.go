package repository

import (
	"api/internal/model/cart"
	"api/internal/model/order"
	"api/internal/model/product"
	"api/internal/model/user"
	"fmt"
	"time"
)

var (
	userSlice         []*user.User
	productSlice      []*product.Product
	orderSlice        []*order.Order
	orderProductSlice []*order.OrderProduct
	cartSlice         []*cart.Cart
	cartProductSlice  []*cart.CartProduct

	logChan chan Item
)

type Item interface {
	SaveItem()
}

func SetItems(itemChannel <-chan Item) {

	for item := range itemChannel {

		switch v := item.(type) {
		case *user.User:
			userSlice = append(userSlice, v)
			AddItem(v)
		case *product.Product:
			productSlice = append(productSlice, v)
			AddItem(v)
		case *cart.Cart:
			cartSlice = append(cartSlice, v)
			AddItem(v)
		case *cart.CartProduct:
			cartProductSlice = append(cartProductSlice, v)
			AddItem(v)
		case *order.Order:
			orderSlice = append(orderSlice, v)
			AddItem(v)
		case *order.OrderProduct:
			orderProductSlice = append(orderProductSlice, v)
			AddItem(v)
		default:
			fmt.Println("Error in repository.SetItems: undefined type of item")
		}

	}

	close(logChan)
}

func AddItem(item Item) {
	item.SaveItem()
	logChan <- item
	// fmt.Print("AddItem: ")
	// fmt.Println(item)
}

func Logger() {

	logChan = make(chan Item)

	oldUserCount := 0
	oldProductCount := 0
	oldCartCount := 0
	oldCartProductCount := 0
	oldOrderCount := 0
	oldOrderProductCount := 0

	fmt.Println("Logger begin working....")

	for item := range logChan {

		time.Sleep(200 * time.Millisecond)

		if len(userSlice) != oldUserCount ||
			len(productSlice) != oldProductCount ||
			len(cartSlice) != oldCartCount ||
			len(cartProductSlice) != oldCartProductCount ||
			len(orderSlice) != oldOrderCount ||
			len(orderProductSlice) != oldOrderProductCount {

			fmt.Println("Log for:")
			fmt.Println(item)

			oldUserCount = len(userSlice)
			oldProductCount = len(productSlice)
			oldCartCount = len(cartSlice)
			oldCartProductCount = len(cartProductSlice)
			oldOrderCount = len(orderSlice)
			oldOrderProductCount = len(orderProductSlice)
		}
	}
}
