package apiv1

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/C22-PS350/backend-rawati/internal/utils"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

// @Summary      get all exercise activity
// @Description  get all exercise activity
// @Tags         activity (exercise)
// @Param        user_id  path   int     true  "User ID"
// @Param        date     query  string  true  "Date (ISO 8601 - date only)"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=[]models.ExerciseActivityResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      403  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /users/{user_id}/exercises [get]
func (h *Handler) GetAllExerciseActivity(w http.ResponseWriter, r *http.Request) {
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

	resp := make([]models.ExerciseActivityResponse, 0)
	if err := h.DB.Table("exercise_per_day").Where("user_id = ? AND exercise_date = ?", userID, date).Find(&resp).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondOK(w, &resp)
}

// @Summary      get exercise activity
// @Description  get exercise activity
// @Tags         activity (exercise)
// @Param        user_id      path  int  true  "User ID"
// @Param        exercise_id  path  int  true  "Exercise Activity ID"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=models.ExerciseActivityResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      403  {object}  utils.JsonErr
// @Failure      404  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /users/{user_id}/exercises/{exercise_id} [get]
func (h *Handler) GetExerciseActivity(w http.ResponseWriter, r *http.Request) {
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

	exerciseID, err := strconv.ParseUint(chi.URLParam(r, "exercise_id"), 10, 64)
	if err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("invalid url parameter form"))
		return
	}

	var resp models.ExerciseActivityResponse
	if err := h.DB.Table("exercise_per_day").Where("exercise_activity_id = ?", exerciseID).First(&resp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondErr(w, http.StatusNotFound, errors.New("record not found"))
			return
		}

		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondOK(w, &resp)
}

// @Summary      create exercise activity
// @Description  create exercise activity
// @Tags         activity (exercise)
// @Accept       json
// @Param        payload  body  models.ExerciseActivityRequest  true  "request body"
// @Param        user_id  path  int                             true  "User ID"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=models.ExerciseActivityCreateResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      403  {object}  utils.JsonErr
// @Failure      404  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /users/{user_id}/exercises [post]
func (h *Handler) CreateExerciseActivity(w http.ResponseWriter, r *http.Request) {
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

	var req models.ExerciseActivityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error decoding request: %s", err))
		return
	}

	if err := h.V.Struct(&req); err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("request body validation error"))
		return
	}

	var userData models.ExerciseActivityUserData
	if err := h.DB.Table("user_profile").Where("user_id = ?", userID).First(&userData).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondErr(w, http.StatusNotFound, errors.New("record not found"))
			return
		}

		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	if err := h.V.Struct(&userData); err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("missing required user profile data"))
		return
	}

	var gender int
	switch userData.Gender {
	case "L":
		gender = 1
	case "P":
		gender = 0
	}

	age := int(math.Floor(time.Since(userData.BirthDate).Hours() / (365 * 24)))
	predictReq := models.ExerciseActivityPredictRequest{
		Gender:    gender,
		Age:       age,
		Weight:    userData.Weight,
		Height:    userData.Height,
		Duration:  req.Duration,
		HeartRate: req.HeartRate,
		BodyTemp:  req.BodyTemp,
	}

	reqBody := &bytes.Buffer{}
	if err := json.NewEncoder(reqBody).Encode(&predictReq); err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	var predictResp *http.Response
	if h.Environment == utils.Testing {
		predictResp = &http.Response{
			Body:       io.NopCloser(strings.NewReader(`{"calories": 200}`)),
			StatusCode: http.StatusOK,
		}
	} else {
		predictResp, err = http.Post(h.ModelAPIUrl, "application/json", io.NopCloser(reqBody))
		if err != nil {
			utils.RespondErr(w, http.StatusInternalServerError, err)
			return
		}
	}
	defer predictResp.Body.Close()

	if predictResp.StatusCode != http.StatusOK {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	var predictResult models.ExerciseActivityPredictResult
	if err := json.NewDecoder(predictResp.Body).Decode(&predictResult); err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error decoding prediction resp: %s", err))
		return
	}

	req.UserID = userID
	req.ExerciseDate = time.Now()
	req.Calories = predictResult.Calories
	if err := h.DB.Table("exercise_per_day").Create(&req).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	resp := models.ExerciseActivityCreateResponse{
		ExerciseActivityID: req.ExerciseActivityID,
		UserID:             userID,
		Calories:           req.Calories,
	}
	utils.RespondOK(w, &resp)
}
