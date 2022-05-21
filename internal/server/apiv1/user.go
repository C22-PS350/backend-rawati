package apiv1

import (
	"encoding/json"
	"net/http"

	"github.com/farryl/project-mars/internal/models"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if req.Name == "" || req.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.DB.Create(&req).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
