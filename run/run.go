package run

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Id       int    `json:"id"`
	Code     string `json:"code"`
	Language string `json:"language"`
	Image    string `json:"image"`
}

type Handler struct {
	Service *Service
}

func (h *Handler) Run(w http.ResponseWriter, r *http.Request) {
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	RespondJSON(w, nil, http.StatusOK)

	err = h.Service.Run(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
