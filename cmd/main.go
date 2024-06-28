package main

import (
	"github.com/prateek041/ecom-go/cmd/api"
)

func main() {
	server := api.NewApiServer(":9090")
	server.Start()
}
