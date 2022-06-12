package apiv1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/C22-PS350/backend-rawati/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// @Summary      get user info
// @Description  get user account info
// @Tags         user
// @Param        user_id  path  int  true  "User ID"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=models.GetUserResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      403  {object}  utils.JsonErr
// @Failure      404  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /users/{user_id} [get]
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseUint(chi.URLParam(r, "user_id"), 10, 64)
	if err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("invalid url parameter form"))
		return
	}

	if h.Environment != utils.Testing {
		if userID != getUserContext(r) {
			utils.RespondErr(w, http.StatusForbidden, errors.New("invalid operation: unmatched user id"))
			return
		}
	}

	var resp models.GetUserResponse
	if err := h.DB.Table("users").First(&resp, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondErr(w, http.StatusNotFound, errors.New("record not found"))
		}

		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondOK(w, &resp)
}

// @Summary      update user info
// @Description  update user account info
// @Tags         user
// @Accept       json
// @Param        user_id  path  int                       true  "User ID"
// @Param        payload  body  models.UpdateUserRequest  true  "request body"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=models.UpdateUserResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      403  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /users/{user_id} [put]
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseUint(chi.URLParam(r, "user_id"), 10, 64)
	if err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("invalid url parameter form"))
		return
	}

	if h.Environment != utils.Testing {
		if userID != getUserContext(r) {
			utils.RespondErr(w, http.StatusForbidden, errors.New("invalid operation: unmatched user id"))
			return
		}
	}

	var req models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error decoding request: %s", err))
		return
	}

	if err := h.V.Struct(&req); err != nil {
		utils.RespondErr(w, http.StatusBadRequest, err)
		return
	}

	req.UserID = userID
	if err := h.DB.Table("users").Save(&req).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			utils.RespondErr(w, http.StatusBadRequest, errors.New("duplicated entry: username or email"))
			return
		}

		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error inserting data: %s", err))
		return
	}

	resp := models.UpdateUserResponse{UserID: req.UserID}
	utils.RespondOK(w, &resp)
}

// @Summary      update user password
// @Description  update authenticated user password
// @Tags         user
// @Accept       json
// @Param        user_id  path  int                          true  "User ID"
// @Param        payload  body  models.UpdateUserPwdRequest  true  "request body"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=models.UpdateUserResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      403  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /users/{user_id}/update-password [put]
func (h *Handler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseUint(chi.URLParam(r, "user_id"), 10, 64)
	if err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("invalid url parameter form"))
		return
	}

	if h.Environment != utils.Testing {
		if userID != getUserContext(r) {
			utils.RespondErr(w, http.StatusForbidden, errors.New("invalid operation: unmatched user id"))
			return
		}
	}

	var req models.UpdateUserPwdRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error decoding request: %s", err))
		return
	}

	if err := h.V.Struct(&req); err != nil {
		utils.RespondErr(w, http.StatusBadRequest, err)
		return
	}

	var data models.UpdateUserPwdData
	if err := h.DB.Table("users").Select("password").Where("user_id = ?", userID).Scan(&data).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	if err := h.V.Struct(&data); err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("invalid operation: user sign up using third party account"))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.OldPassword)); err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("invalid old password"))
		return
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	if err := h.DB.Table("users").Where("user_id = ?", userID).Update("password", string(hashedPwd)).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	resp := models.UpdateUserResponse{UserID: userID}
	utils.RespondOK(w, &resp)
}

// @Summary      delete user
// @Description  delete user account
// @Tags         user
// @Param        user_id  path  int  true  "User ID"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=models.UpdateUserResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      403  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /users/{user_id} [delete]
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseUint(chi.URLParam(r, "user_id"), 10, 64)
	if err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("invalid url parameter form"))
		return
	}

	var t string
	if h.Environment != utils.Testing {
		if userID != getUserContext(r) {
			utils.RespondErr(w, http.StatusForbidden, errors.New("invalid operation: unmatched user id"))
			return
		}

		if err := h.DB.Raw(authFindTokenByUser, userID).Scan(&t).Error; err != nil {
			utils.RespondErr(w, http.StatusInternalServerError, err)
			return
		}
	}

	if err := h.DB.Table("users").Where("user_id = ?", userID).Delete(&struct{}{}).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	if h.Environment != utils.Testing {
		h.C.Delete(t)
	}

	resp := models.UpdateUserResponse{UserID: userID}
	utils.RespondOK(w, &resp)
}
