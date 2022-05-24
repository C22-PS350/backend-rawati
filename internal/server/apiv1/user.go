package apiv1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/C22-PS350/backend-rawati/internal/util"
)

// @Summary      create user
// @Description  create a user account
// @Accept       json
// @Param        payload  body  models.UserRequest  true  "request body"
// @Produce      json
// @Success      200  {object}  util.JsonOK{data=models.UserResponse}
// @Failure      400  {object}  util.JsonErr
// @Failure      500  {object}  util.JsonErr
// @Router       /users [post]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error decoding request: %s", err))
		return
	}

	if req.Name == "" || req.Password == "" {
		util.RespondErr(w, http.StatusBadRequest, errors.New("name or password can't be empty"))
		return
	}

	if err := h.DB.Create(&req).Error; err != nil {
		util.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error inserting data: %s", err))
		return
	}

	util.RespondOK(w, &req)
}
