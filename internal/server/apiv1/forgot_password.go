package apiv1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"cloud.google.com/go/pubsub"
	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/C22-PS350/backend-rawati/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

var (
	forgotPasswordFindUser = `
		SELECT user_id, username, email
		FROM users
		WHERE email = ?
	`
)

// @Summary      forgot password
// @Description  update unauthenticated user password
// @Tags         auth
// @Accept       json
// @Param        payload  body  models.ForgotPasswordRequest  true  "request body"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=models.ForgotPasswordResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /auth/forgot-password [put]
func (h *Handler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var req models.ForgotPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error decoding request: %s", err))
		return
	}

	if err := h.V.Struct(&req); err != nil {
		utils.RespondErr(w, http.StatusBadRequest, err)
		return
	}

	var userData models.ForgotPasswordData
	if err := h.DB.Raw(forgotPasswordFindUser, req.Email).Scan(&userData).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	if userData.UserID == 0 {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("user with requested email not found"))
		return
	}

	newPwd := generateToken(15)
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(newPwd), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error hashing password: %s", err))
		return
	}

	if err := h.DB.Table("users").Where("user_id = ?", userData.UserID).Update("password", string(hashedPwd)).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	if h.Environment != utils.Remote {
		utils.RespondOK(w, newPwd)
		return
	}

	topic := h.GcpClients.PubSub.Topic("mailbox")
	ok, err := topic.Exists(context.Background())
	if err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}
	if !ok {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}
	defer topic.Stop()

	fpmud := models.ForgotPasswordMessage{
		Username: userData.Username,
		Email:    userData.Email,
		Password: newPwd,
	}

	m, err := json.Marshal(&fpmud)
	if err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error marshaling pub/sub message: %s", err))
		return
	}

	res := topic.Publish(r.Context(), &pubsub.Message{
		Data: m,
		Attributes: map[string]string{
			"type": "forgotpassword",
		},
	})

	_, err = res.Get(r.Context())
	if err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	resp := models.ForgotPasswordResponse{Message: "forgot password request success"}
	utils.RespondOK(w, &resp)
}
