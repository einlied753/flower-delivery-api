package main

import (
	"api/internal/repository"
	"api/internal/service"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		repository.Logger()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		service.CreateItems()
		wg.Done()
	}()

	wg.Wait()
}
