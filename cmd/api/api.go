package api

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prateek041/ecom-go/services/test"
)

type ApiServer struct {
	server    *http.Server
	l         *log.Logger
	apiRouter *mux.Router
}

func NewApiServer(addr string, l *log.Logger) *ApiServer {
	server := mux.NewRouter()
	apiRouter := server.PathPrefix("/api/v1").Subrouter()

	// TODO: Define it more with options like read, write and idle timeouts
	return &ApiServer{
		server: &http.Server{
			Addr:    addr,
			Handler: apiRouter,
		},
		l:         l,
		apiRouter: apiRouter,
	}
}

func (s *ApiServer) Run() error {
	// Register handlers and database connections
	s.registerRoutes()
	return s.server.ListenAndServe()
}

func (s *ApiServer) ShutDown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *ApiServer) registerRoutes() {
	testHandler := test.NewTestHandler(s.l)
	testHandler.RegisterRoutes(s.apiRouter)
}
