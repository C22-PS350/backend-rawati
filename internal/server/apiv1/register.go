package apiv1

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/C22-PS350/backend-rawati/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// @Summary      register user
// @Description  register a user account
// @Accept       json
// @Param        payload  body  models.RegisterRequest true  "request body"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=models.RegisterResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /auth/register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error decoding request: %s", err))
		return
	}

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("request body validation error"))
		return
	}

	pwd := *(req.Password)
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error hashing password: %s", err))
		return
	}

	hashedPwdStr := string(hashedPwd)
	req.Password = &hashedPwdStr
	if err := h.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("users").Create(&req).Error; err != nil {
			return err
		}

		userToken := models.RegisterUserToken{
			UserID:    req.UserID,
			Token:     generateToken(20),
			CreatedAt: time.Now(),
		}

		if err := tx.Table("user_token").Create(&userToken).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			utils.RespondErr(w, http.StatusBadRequest, errors.New("duplicated entry: username or email"))
			return
		}

		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error inserting data: %s", err))
		return
	}

	resp := models.RegisterResponse{UserID: req.UserID}
	utils.RespondOK(w, &resp)
}

func generateToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
