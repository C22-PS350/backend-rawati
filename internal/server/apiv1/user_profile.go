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
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// @Summary      get user profile info
// @Description  get user profile info
// @Tags         user profile
// @Param        user_id path int true "User ID"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=models.GetUserProfileResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      403  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /users/{user_id}/profile [get]
func (h *Handler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
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

	var resp models.GetUserProfileResponse
	if err := h.DB.Table("user_profile").Where("user_id = ?", userID).First(&resp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondErr(w, http.StatusNotFound, errors.New("record not found"))
			return
		}

		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondOK(w, &resp)
}

// @Summary      update user profile info
// @Description  update user profile info
// @Tags         user profile
// @Accept       json
// @Param        user_id path int true "User ID"
// @Param        payload  body  models.UserProfileRequest true "request body"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=models.UserProfileResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      403  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /users/{user_id}/profile [put]
func (h *Handler) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
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

	var req models.UserProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error decoding request: %s", err))
		return
	}

	var userData models.GetUserProfileResponse
	if err := h.DB.Table("user_profile").Clauses(clause.Returning{Columns: []clause.Column{{Name: "profile_id"}}}).Where("user_id = ?", userID).Updates(req).Scan(&userData).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	resp := models.UserProfileResponse{ProfileID: userData.ProfileID, UserID: userID}
	utils.RespondOK(w, &resp)
}

// func (h *Handler) CreateUserProfile(w http.ResponseWriter, r *http.Request) {
// 	var req models.UserProfileRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error decoding request: %s", err))
// 		return
// 	}

// 	validate := validator.New()
// 	if err := validate.Struct(&req); err != nil {
// 		utils.RespondErr(w, http.StatusBadRequest, errors.New("request body validation error"))
// 		return
// 	}

// 	if h.Environment != utils.Testing {
// 		req.UserID = getUserContext(r)
// 	}

// 	if err := h.DB.Table("user_profile").Create(&req).Error; err != nil {
// 		utils.RespondErr(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	resp := models.UserProfileResponse{ProfileID: req.ProfileID, UserID: req.UserID}
// 	utils.RespondOK(w, &resp)
// }
