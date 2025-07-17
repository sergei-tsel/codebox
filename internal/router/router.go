package router

import (
	"codebox/internal/service"
	"codebox/internal/utils"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/ping", pong)

	r.Post("/api/run", run)

	return r
}

func pong(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("pong"))
}

func run(w http.ResponseWriter, r *http.Request) {
	var req service.RunRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondJSON(w, nil, http.StatusOK)

	err = service.Run(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
