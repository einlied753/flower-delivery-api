package repository

import (
	"api/internal/model/cart"
	"log"
	"strconv"
	"sync"
)

var (
	cartSlice []*cart.Cart
	cartMutex sync.RWMutex
)

const cartInputFile string = "./input_files/carts.csv"
const cartOutputFile string = "./output_files/carts.csv"

func GetCarts() []*cart.Cart {
	cartMutex.RLock()
	defer cartMutex.RUnlock()
	return cartSlice
}

func ReadNewCarts(itemChan chan<- Item) {
	carts, err := readFile(cartInputFile)
	if err != nil {
		log.Fatalf("An error occurred while reading the new carts from file: %v", err)
	}

	for i, c := range carts {
		userId, err := strconv.Atoi(c[1])
		if err != nil {
			log.Fatalf("An error occurred while converting userId of the new cart: %v", err)
		}
		orderId, err := strconv.Atoi(c[1])
		if err != nil {
			log.Fatalf("An error occurred while converting orderId of the new cart: %v", err)
		}

		itemChan <- cart.NewCart(i+1, userId, orderId)
	}
}

func SaveCart(cart *cart.Cart) {
	cartMutex.Lock()
	defer cartMutex.Unlock()
	cartSlice = append(cartSlice, cart)
	SaveCartToFile(cart)
}

func SaveCartToFile(cart *cart.Cart) {

	strCart := []string{
		strconv.Itoa(cart.GetId()),
		strconv.Itoa(cart.GetUserId()),
		strconv.Itoa(cart.GetOrderId()),
		cart.GetStringCartStatus(),
		cart.GetCreated(),
	}

	writeFile(cartOutputFile, strCart)
}
