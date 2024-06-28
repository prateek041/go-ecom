package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	server http.Server
}

func NewApiServer(addr string) *ApiServer {
	server := mux.NewRouter()
	apiRouter := server.PathPrefix("api/v1").Subrouter()

	// TODO: Define it more with options like read, write and idle timeouts
	return &ApiServer{
		server: http.Server{
			Addr:    addr,
			Handler: apiRouter,
		},
	}
}

func (s *ApiServer) Run() error {
	// Register handlers and database connections
	return s.server.ListenAndServe()
}

func (s *ApiServer) ShutDown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
