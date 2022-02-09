package handler

import (
	"net/http"
	"path"

	"github.com/alexm24/cache-img/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services}
}

func (h *Handler) InitRoutes(basePath string) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route(path.Join("/", basePath), func(r chi.Router) {
		r.Use(middleware.NoCache)
		r.Mount("/", h.rHandler(r))
	})

	return router
}

func (h *Handler) rHandler(r chi.Router) http.Handler {
	r.Get("/avatar/{code}", h.getAvatar)
	return r
}
