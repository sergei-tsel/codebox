package handler

import (
	"codebox/service"
	"codebox/utils"
	"encoding/json"
	"net/http"
)

type RunHandler struct {
	RunService *service.RunService
}

func (h *RunHandler) Run(w http.ResponseWriter, r *http.Request) {
	var req service.RunRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondJSON(w, nil, http.StatusOK)

	err = h.RunService.Run(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
