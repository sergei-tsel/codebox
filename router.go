package main

import (
	"codebox/run"
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	h := run.Handler{
		Service: &run.Service{
			Repo: &run.Repo{},
		},
	}

	r.Get("/ping", Pong)

	r.Post("/api/run", h.Run)

	return r
}
