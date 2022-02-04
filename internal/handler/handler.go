package handler

import (
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route(path.Join("/", cfg, BasePath), func(r chi.Router) {
		r.Use(middleware.NoCache)
		r.Mount("/")
	})
	router.Get("/avatar/{code}", h.getAvatar)
	return router
}
