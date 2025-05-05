package main

import (
	"api/internal/service"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		service.ItemSavingLogging()
		wg.Done()
	}()

	go func() {
		service.CreateItems()
		wg.Done()
	}()

	wg.Wait()
}
