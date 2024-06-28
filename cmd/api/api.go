package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
	// TODO: Add database string here as well
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{
		addr: addr,
	}
}

func (s *ApiServer) Start() {
	server := mux.NewRouter()
	apiRouter := server.PathPrefix("api/v1").Subrouter()
	fmt.Printf("Starting server on addr %s", s.addr)
	err := http.ListenAndServe(s.addr, apiRouter)
	if err != nil {
		log.Fatal("Error starting the server", err)
	}
}
