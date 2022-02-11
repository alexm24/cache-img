package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) SingUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test1")
}

func (h *Handler) SingIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test2")
}
