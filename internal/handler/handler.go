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
		r.Get("/avatar/{code}", h.getAvatar)
		r.Post("/auth/sing-up", h.SingUp)
		r.Post("/auth/sing-in", h.SingIn)
	})

	return router
}
