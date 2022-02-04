package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) getAvatar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
}
