package service

import (
	"api/internal/model/cart"
	"api/internal/model/order"
	"api/internal/model/product"
	"api/internal/model/user"
	"api/internal/repository"
	"fmt"
	"slices"
	"sync"
	"time"
)

func ItemSavingLogging() {

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

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

	fmt.Println("Logger started working...")

	for {
		mu.Lock()
		newUsers := repository.GetUsers()
		mu.Unlock()

		if len(newUsers) != oldUsersCount {

			wg.Add(2)
			go func() {
				for user := range userChan {
					user.SaveItemLog()
				}
				wg.Done()
			}()
			go func() {
				for _, u := range newUsers {
					if !slices.Contains(addedUsers, u) {
						userChan <- u
						addedUsers = append(addedUsers, u)
					}
				}
				wg.Done()
			}()
			oldUsersCount = len(newUsers)
		}

		mu.Lock()
		newProducts := repository.GetProducts()
		mu.Unlock()

		if len(newProducts) != oldProductCount {

			wg.Add(2)
			go func() {
				for product := range productChan {
					product.SaveItemLog()
				}
				wg.Done()
			}()
			go func() {
				for _, p := range newProducts {
					if !slices.Contains(addedProducts, p) {
						productChan <- p
						addedProducts = append(addedProducts, p)
					}
				}
				wg.Done()
			}()
			oldProductCount = len(newProducts)
		}

		mu.Lock()
		newCarts := repository.GetCarts()
		mu.Unlock()

		if len(newCarts) != oldCartCount {

			wg.Add(2)
			go func() {
				for cart := range cartChan {
					cart.SaveItemLog()
				}
				wg.Done()
			}()
			go func() {
				for _, c := range newCarts {
					if !slices.Contains(addedCarts, c) {
						cartChan <- c
						addedCarts = append(addedCarts, c)
					}
				}
				wg.Done()
			}()
			oldCartCount = len(newCarts)
		}

		mu.Lock()
		newCartProducts := repository.GetCartProducts()
		mu.Unlock()

		if len(newCartProducts) != oldCartProductCount {

			wg.Add(2)
			go func() {
				for cartProduct := range cartProductChan {
					cartProduct.SaveItemLog()
				}
				wg.Done()
			}()
			go func() {
				for _, cp := range newCartProducts {
					if !slices.Contains(addedCartProducts, cp) {
						cartProductChan <- cp
						addedCartProducts = append(addedCartProducts, cp)
					}
				}
				wg.Done()
			}()
			oldCartProductCount = len(newCartProducts)
		}

		mu.Lock()
		newOrders := repository.GetOrders()
		mu.Unlock()

		if len(newOrders) != oldOrderCount {

			wg.Add(2)
			go func() {
				for order := range orderChan {
					order.SaveItemLog()
				}
				wg.Done()
			}()
			go func() {
				for _, o := range newOrders {
					if !slices.Contains(addedOrders, o) {
						orderChan <- o
						addedOrders = append(addedOrders, o)
					}
				}
				wg.Done()
			}()
			oldOrderCount = len(newOrders)
		}

		mu.Lock()
		newOrderProducts := repository.GetOrderProducts()
		mu.Unlock()

		if len(newOrderProducts) != oldOrderProductCount {

			wg.Add(2)
			go func() {
				for orderProduct := range orderProductChan {
					orderProduct.SaveItemLog()
				}
				wg.Done()
			}()
			go func() {
				for _, op := range newOrderProducts {
					if !slices.Contains(addedOrderProducts, op) {
						orderProductChan <- op
						addedOrderProducts = append(addedOrderProducts, op)
					}
				}
				wg.Done()
			}()
			oldOrderProductCount = len(newOrderProducts)
		}

		time.Sleep(200 * time.Millisecond)
	}
}
