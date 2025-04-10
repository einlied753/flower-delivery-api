package repository

import (
	"api/internal/model/cart"
	"api/internal/model/order"
	"api/internal/model/product"
	"api/internal/model/user"
	"fmt"
	"strconv"
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
	SaveItem()
}

func SetItems(items []Item) {
	for _, item := range items {
		switch v := item.(type) {
		case *user.User:
			userSlice = append(userSlice, v)
			item.SaveItem()
		case *product.Product:
			productSlice = append(productSlice, v)
			item.SaveItem()
		case *cart.Cart:
			cartSlice = append(cartSlice, v)
			item.SaveItem()
		case *cart.CartProduct:
			cartProductSlice = append(cartProductSlice, v)
			item.SaveItem()
		case *order.Order:
			orderSlice = append(orderSlice, v)
			item.SaveItem()
		case *order.OrderProduct:
			orderProductSlice = append(orderProductSlice, v)
			item.SaveItem()
		default:
			fmt.Println("Error in repository.GetItems: undefined type of item")
		}
	}

	fmt.Println("Users - " + strconv.Itoa(len(userSlice)))
	fmt.Println("Products - " + strconv.Itoa(len(productSlice)))
	fmt.Println("Carts - " + strconv.Itoa(len(cartSlice)))
	fmt.Println("CartProducts - " + strconv.Itoa(len(cartProductSlice)))
	fmt.Println("Orders - " + strconv.Itoa(len(orderSlice)))
	fmt.Println("OrderProducts - " + strconv.Itoa(len(orderProductSlice)))
}
