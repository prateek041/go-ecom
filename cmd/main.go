package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/prateek041/ecom-go/cmd/api"
)

const addr = ":9090"

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	server := api.NewApiServer(addr)

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
	err := server.ShutDown(tc)
	logger.Fatal("Error shutting down server", err)
}
