package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/prateek041/ecom-go/cmd/api"
	"github.com/prateek041/ecom-go/configs"
	"github.com/prateek041/ecom-go/db"
)

const addr = ":9090"

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	storageProvider, err := db.NewStorage(configs.ENV.DBUri)
	if err != nil {
		logger.Fatalf("error initializing storage: %v", err)
	}
	server := api.NewApiServer(addr, logger, storageProvider.Client)

	go func() {
		logger.Printf("Starting Server on addr %s", addr)
		err := server.Run()
		if err != nil {
			logger.Fatal("Shutting down server", err)
		}
	}()

	killChannel := make(chan os.Signal)
	signal.Notify(killChannel, os.Kill)
	signal.Notify(killChannel, os.Interrupt)

	sig := <-killChannel
	logger.Println("Recieved signal, Graceful Shutdown", sig)

	// ignoring cancel function
	tc, _ := context.WithTimeout(context.Background(), time.Second*30)
	err = server.ShutDown(tc)
	logger.Fatal("Error shutting down server", err)
}
