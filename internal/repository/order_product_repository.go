package repository

import (
	"api/internal/model/order"
	"log"
	"strconv"
	"sync"
)

var (
	orderProductSlice []*order.OrderProduct
	orderProductMutex sync.RWMutex
)

const orderProductInputFile string = "./input_files/order_products.csv"
const orderProductOutputFile string = "./output_files/order_products.csv"

func GetOrderProducts() []*order.OrderProduct {
	orderProductMutex.RLock()
	defer orderProductMutex.RUnlock()
	return orderProductSlice
}

func ReadNewOrderProducts(itemChan chan<- Item) {
	orderProducts, err := readFile(orderProductInputFile)
	if err != nil {
		log.Fatalf("An error occurred while reading the new order products from file: %v", err)
	}

	for i, op := range orderProducts {
		orderId, err := strconv.Atoi(op[1])
		if err != nil {
			log.Fatalf("An error occurred while converting orderId of the new order product: %v", err)
		}
		productId, err := strconv.Atoi(op[2])
		if err != nil {
			log.Fatalf("An error occurred while converting productId of the new order product: %v", err)
		}
		productCount, err := strconv.Atoi(op[3])
		if err != nil {
			log.Fatalf("An error occurred while converting productCount of the new order product: %v", err)
		}

		itemChan <- order.NewOrderProduct(i+1, orderId, productId, productCount)
	}
}

func SaveOrderProduct(orderProduct *order.OrderProduct) {
	orderProductMutex.Lock()
	defer orderProductMutex.Unlock()
	orderProductSlice = append(orderProductSlice, orderProduct)
	SaveOrderProductToFile(orderProduct)
}

func SaveOrderProductToFile(orderProduct *order.OrderProduct) {

	strOrderProduct := []string{
		strconv.Itoa(orderProduct.GetId()),
		strconv.Itoa(orderProduct.GetOrderId()),
		strconv.Itoa(orderProduct.GetProductId()),
		strconv.Itoa(orderProduct.GetProductCount()),
	}

	writeFile(orderProductOutputFile, strOrderProduct)
}
