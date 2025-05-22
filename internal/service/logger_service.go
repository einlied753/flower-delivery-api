package service

import (
	"api/internal/model/cart"
	"api/internal/model/order"
	"api/internal/model/product"
	"api/internal/model/user"
	"api/internal/repository"
	"context"
	"fmt"
	"slices"
	"time"
)

func ItemSavingLogging(ctx context.Context) {

	oldUsersCount := 0
	oldProductCount := 0
	oldCartCount := 0
	oldCartProductCount := 0
	oldOrderCount := 0
	oldOrderProductCount := 0

	var addedUsers []*user.User
	var addedProducts []*product.Product
	var addedCarts []*cart.Cart
	var addedCartProducts []*cart.CartProduct
	var addedOrders []*order.Order
	var addedOrderProducts []*order.OrderProduct

	userChan := make(chan *user.User)
	productChan := make(chan *product.Product)
	cartChan := make(chan *cart.Cart)
	cartProductChan := make(chan *cart.CartProduct)
	orderChan := make(chan *order.Order)
	orderProductChan := make(chan *order.OrderProduct)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context done in ItemSavingLogging()")
			return
		default:

			newUsers := repository.GetUsers()

			if len(newUsers) != oldUsersCount {

				go func() {
					for user := range userChan {
						user.SaveItemLog()
					}

				}()
				go func() {
					for _, u := range newUsers {
						if !slices.Contains(addedUsers, u) {
							userChan <- u
							addedUsers = append(addedUsers, u)
						}
					}

				}()
				oldUsersCount = len(newUsers)
			}

			newProducts := repository.GetProducts()

			if len(newProducts) != oldProductCount {

				go func() {
					for product := range productChan {
						product.SaveItemLog()
					}
				}()
				go func() {
					for _, p := range newProducts {
						if !slices.Contains(addedProducts, p) {
							productChan <- p
							addedProducts = append(addedProducts, p)
						}
					}
				}()
				oldProductCount = len(newProducts)
			}

			newCarts := repository.GetCarts()

			if len(newCarts) != oldCartCount {

				go func() {
					for cart := range cartChan {
						cart.SaveItemLog()
					}
				}()
				go func() {
					for _, c := range newCarts {
						if !slices.Contains(addedCarts, c) {
							cartChan <- c
							addedCarts = append(addedCarts, c)
						}
					}
				}()
				oldCartCount = len(newCarts)
			}

			newCartProducts := repository.GetCartProducts()

			if len(newCartProducts) != oldCartProductCount {

				go func() {
					for cartProduct := range cartProductChan {
						cartProduct.SaveItemLog()
					}
				}()
				go func() {
					for _, cp := range newCartProducts {
						if !slices.Contains(addedCartProducts, cp) {
							cartProductChan <- cp
							addedCartProducts = append(addedCartProducts, cp)
						}
					}
				}()
				oldCartProductCount = len(newCartProducts)
			}

			newOrders := repository.GetOrders()

			if len(newOrders) != oldOrderCount {

				go func() {
					for order := range orderChan {
						order.SaveItemLog()
					}
				}()
				go func() {
					for _, o := range newOrders {
						if !slices.Contains(addedOrders, o) {
							orderChan <- o
							addedOrders = append(addedOrders, o)
						}
					}
				}()
				oldOrderCount = len(newOrders)
			}

			newOrderProducts := repository.GetOrderProducts()

			if len(newOrderProducts) != oldOrderProductCount {

				go func() {
					for orderProduct := range orderProductChan {
						orderProduct.SaveItemLog()
					}
				}()
				go func() {
					for _, op := range newOrderProducts {
						if !slices.Contains(addedOrderProducts, op) {
							orderProductChan <- op
							addedOrderProducts = append(addedOrderProducts, op)
						}
					}
				}()
				oldOrderProductCount = len(newOrderProducts)
			}

			time.Sleep(200 * time.Millisecond)
		}
	}
}
