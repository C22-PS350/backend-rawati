package apiv1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/C22-PS350/backend-rawati/internal/utils"
)

// @Summary      create user
// @Description  create a user account
// @Accept       json
// @Param        payload  body  models.User  true  "request body"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=models.UserResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /users [post]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error decoding request: %s", err))
		return
	}

	if req.Name == nil || req.Password == nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("name or password can't be empty"))
		return
	}

	if err := h.DB.Table("users").Create(&req).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error inserting data: %s", err))
		return
	}

	utils.RespondOK(w, &req)
}
