package apiv1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/C22-PS350/backend-rawati/internal/utils"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var (
	loginFindUserByIdentifier = `
		SELECT user_id, password, token
		FROM users u
		JOIN user_token ut USING (user_id)
		WHERE u.username = ? OR u.email = ?
	`
)

// @Summary      login user
// @Description  log a user account in
// @Tags         auth
// @Accept       json
// @Param        payload  body  models.LoginRequest  true  "request body"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=models.LoginResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /auth/login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error decoding request: %s", err))
		return
	}

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("request body validation error"))
		return
	}

	var data models.LoginData
	if err := h.DB.Raw(loginFindUserByIdentifier, req.Identifier, req.Identifier).Scan(&data).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	if err := validate.Struct(&data); err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("invalid username, email or password"))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password)); err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("invalid username, email or password"))
		return
	}

	resp := models.LoginResponse{UserID: data.UserID, Token: data.Token}
	utils.RespondOK(w, &resp)
}
