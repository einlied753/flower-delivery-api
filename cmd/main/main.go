package main

import (
	"api/internal/service"
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(2)
	go func() {
		service.ItemSavingLogging(ctx)
		wg.Done()
	}()
	go func() {
		service.CreateItems(ctx)
		wg.Done()
	}()

	<-stopChan
	cancel()

	wg.Wait()
}
