package user

import (
	"log"
	"net/http"
)

type UserHandler struct {
	l *log.Logger
}

func NewUserHandler(l *log.Logger) *UserHandler {
	return &UserHandler{
		l: l,
	}
}

// functionalities that I want to be present with the user.

func (uh *UserHandler) login(rw http.ResponseWriter, r *http.Request) {
}

func (uh *UserHandler) signup(rw http.ResponseWriter, r *http.Request) {
}
