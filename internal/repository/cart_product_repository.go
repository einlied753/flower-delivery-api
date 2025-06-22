package repository

import (
	"api/internal/model/cart"
	"log"
	"strconv"
	"sync"
)

var (
	cartProductSlice []*cart.CartProduct
	cartProductMutex sync.RWMutex
)

const cartProductInputFile string = "./input_files/cart_products.csv"
const cartProductOutputFile string = "./output_files/cart_products.csv"

func GetCartProducts() []*cart.CartProduct {
	cartProductMutex.RLock()
	defer cartProductMutex.RUnlock()
	return cartProductSlice
}

func ReadNewCartProducts(itemChan chan<- Item) {
	cartProducts, err := readFile(cartProductInputFile)
	if err != nil {
		log.Fatalf("An error occurred while reading the new cart products from file: %v", err)
	}

	for i, cp := range cartProducts {
		cartId, err := strconv.Atoi(cp[1])
		if err != nil {
			log.Fatalf("An error occurred while converting cartId of the new cart product: %v", err)
		}
		productId, err := strconv.Atoi(cp[2])
		if err != nil {
			log.Fatalf("An error occurred while converting productId of the new cart product: %v", err)
		}
		productCount, err := strconv.Atoi(cp[3])
		if err != nil {
			log.Fatalf("An error occurred while converting productCount of the new cart product: %v", err)
		}

		itemChan <- cart.NewCartProduct(i+1, cartId, productId, productCount)
	}
}

func SaveCartProduct(cartProduct *cart.CartProduct) {
	cartProductMutex.Lock()
	defer cartProductMutex.Unlock()
	cartProductSlice = append(cartProductSlice, cartProduct)
	SaveCartProductToFile(cartProduct)
}

func SaveCartProductToFile(cartProduct *cart.CartProduct) {

	strCartProduct := []string{
		strconv.Itoa(cartProduct.GetId()),
		strconv.Itoa(cartProduct.GetCartId()),
		strconv.Itoa(cartProduct.GetProductId()),
		strconv.Itoa(cartProduct.GetProductCount()),
	}

	writeFile(cartProductOutputFile, strCartProduct)
}
