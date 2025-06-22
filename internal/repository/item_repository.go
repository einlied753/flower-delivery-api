package repository

import (
	"api/internal/model/cart"
	"api/internal/model/order"
	"api/internal/model/product"
	"api/internal/model/user"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Item interface {
	SaveItemLog()
}

func AddItem(item Item) {
	switch v := item.(type) {
	case *user.User:
		SaveUser(v)
	case *product.Product:
		SaveProduct(v)
	case *cart.Cart:
		SaveCart(v)
	case *cart.CartProduct:
		SaveCartProduct(v)
	case *order.Order:
		SaveOrder(v)
	case *order.OrderProduct:
		SaveOrderProduct(v)
	default:
		fmt.Println("Error in AddItem: undefined type of item")
	}
}

func writeFile(fileName string, strItem []string) {

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Unable to open file %s, err: %v", fileName, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(strItem); err != nil {
		log.Fatalf("Cannot write row: %v", err)
	}
}

func readFile(filePath string) ([][]string, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("couldn't open file %s, error: %w", filePath, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("csvReader error: %w", err)
	}

	return lines, nil
}
