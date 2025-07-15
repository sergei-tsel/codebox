package router

import (
	"codebox/handler"
	"codebox/repository"
	"codebox/service"
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	h := handler.RunHandler{
		RunService: &service.RunService{
			Repo: &repository.RunRepo{},
		},
	}

	r.Get("/ping", handler.Pong)

	r.Post("/api/run", h.Run)

	return r
}
