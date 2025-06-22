package repository

import (
	"api/internal/model/product"
	"log"
	"strconv"
	"sync"
)

var (
	productSlice []*product.Product
	productMutex sync.RWMutex
)

const productInputFile string = "./input_files/products.csv"
const productOutputFile string = "./output_files/products.csv"

func GetProducts() []*product.Product {
	productMutex.RLock()
	defer productMutex.RUnlock()
	return productSlice
}

func ReadNewProducts(itemChan chan<- Item) {
	products, err := readFile(productInputFile)
	if err != nil {
		log.Fatalf("An error occurred while reading the new products from file: %v", err)
	}

	for i, p := range products {

		price, err := strconv.ParseFloat(p[2], 32)
		if err != nil {
			log.Fatalf("An error occurred while converting price of the new product: %v", err)
		}

		quantity, err := strconv.Atoi(p[3])
		if err != nil {
			log.Fatalf("An error occurred while converting quantity of the new product: %v", err)
		}

		discount, err := strconv.ParseFloat(p[4], 32)
		if err != nil {
			log.Fatalf("An error occurred while converting discount of the new product: %v", err)
		}

		itemChan <- product.NewProduct(i+1, p[1], float32(price), quantity, float32(discount))
	}
}

func SaveProduct(product *product.Product) {
	productMutex.Lock()
	defer productMutex.Unlock()
	productSlice = append(productSlice, product)
	SaveProductToFile(product)
}

func SaveProductToFile(product *product.Product) {
	strProduct := []string{
		strconv.Itoa(product.GetId()),
		product.GetName(),
		strconv.FormatFloat(float64(product.GetPrice()), 'f', 2, 64),
		strconv.Itoa(product.GetQuantity()),
		strconv.FormatFloat(float64(product.GetDiscount()), 'f', 2, 64),
	}

	writeFile(productOutputFile, strProduct)
}
