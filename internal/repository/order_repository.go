package repository

import (
	"api/internal/model/order"
	"log"
	"strconv"
	"sync"
)

var (
	orderSlice []*order.Order
	orderMutex sync.RWMutex
)

const orderInputFile string = "./input_files/orders.csv"
const orderOutputFile string = "./output_files/orders.csv"

func GetOrders() []*order.Order {
	orderMutex.RLock()
	defer orderMutex.RUnlock()
	return orderSlice
}

func ReadNewOrders(itemchan chan<- Item) {
	orders, err := readFile(orderInputFile)
	if err != nil {
		log.Fatalf("An error occurred while reading the new orders from file: %v", err)
	}

	for i, o := range orders {

		cost, err := strconv.ParseFloat(o[1], 32)
		if err != nil {
			log.Fatalf("An error occurred while converting cost of the new order: %v", err)
		}

		productCount, err := strconv.Atoi(o[2])
		if err != nil {
			log.Fatalf("An error occurred while converting productCount of the new order: %v", err)
		}

		itemchan <- order.NewOrder(i+1, float32(cost), productCount, o[3], o[4], o[5])
	}
}

func SaveOrder(order *order.Order) {
	orderMutex.Lock()
	defer orderMutex.Unlock()
	orderSlice = append(orderSlice, order)
	SaveOrderToFile(order)
}

func SaveOrderToFile(order *order.Order) {

	strOrder := []string{
		strconv.Itoa(order.GetId()),
		strconv.Itoa(order.GetUserId()),
		order.GetStringOrderStatus(),
		strconv.FormatFloat(float64(order.GetCost()), 'f', 2, 64),
		strconv.Itoa(order.GetProductCount()),
		order.GetCreated(),
		order.GetEmail(),
		order.GetAddress(),
		order.GetPhone(),
	}

	writeFile(orderOutputFile, strOrder)
}
