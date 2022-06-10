package apiv1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/C22-PS350/backend-rawati/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// @Summary      get all food activity
// @Description  get all food activity
// @Tags         activity food
// @Param        user_id  path   int     true  "User ID"
// @Param        date     query  string  true  "Date (ISO 8601 - date only)"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=[]models.FoodActivityResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      403  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /users/{user_id}/foods [get]
func (h *Handler) GetAllFoodActivity(w http.ResponseWriter, r *http.Request) {
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

	date := r.URL.Query().Get("date")
	if date == "" {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("date query parameter is missing"))
		return
	}

	_, err = time.Parse("2006-01-02", date)
	if err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("invalid date query parameter value"))
		return
	}

	resp := make([]models.FoodActivityResponse, 0)
	if err := h.DB.Table("food_per_day").Where("user_id = ? AND food_date = ?", userID, date).Find(&resp).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondOK(w, &resp)
}

// @Summary      get food activity
// @Description  get food activity
// @Tags         activity food
// @Param        user_id  path  int  true  "User ID"
// @Param        food_id  path  int  true  "Food Activity ID"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=models.FoodActivityResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      403  {object}  utils.JsonErr
// @Failure      404  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /users/{user_id}/foods/{food_id} [get]
func (h *Handler) GetFoodActivity(w http.ResponseWriter, r *http.Request) {
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

	foodID, err := strconv.ParseUint(chi.URLParam(r, "food_id"), 10, 64)
	if err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("invalid url parameter form"))
		return
	}

	var resp models.FoodActivityResponse
	if err := h.DB.Table("food_per_day").Where("food_activity_id = ?", foodID).First(&resp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondErr(w, http.StatusNotFound, errors.New("record not found"))
			return
		}

		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondOK(w, &resp)
}

// @Summary      create food activity
// @Description  create food activity
// @Tags         activity food
// @Accept       json
// @Param        payload  body  models.FoodActivityRequest  true  "request body"
// @Param        user_id  path  int                         true  "User ID"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=models.FoodActivityCreateResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      403  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /users/{user_id}/foods [post]
func (h *Handler) CreateFoodActivity(w http.ResponseWriter, r *http.Request) {
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

	var req models.FoodActivityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error decoding request: %s", err))
		return
	}

	req.FoodDate = time.Now()
	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("request body validation error"))
		return
	}

	req.UserID = userID
	if err := h.DB.Table("food_per_day").Create(&req).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	resp := models.FoodActivityCreateResponse{FoodActivityID: req.FoodActivityID, UserID: userID}
	utils.RespondOK(w, &resp)
}
