package test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TestHandler struct {
	l *log.Logger
}

func NewTestHandler(l *log.Logger) *TestHandler {
	return &TestHandler{
		l: l,
	}
}

func (t *TestHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/ping", t.ping)
}

func (t *TestHandler) ping(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "pong")
}
