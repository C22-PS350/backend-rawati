package apiv1

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/C22-PS350/backend-rawati/internal/utils"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

// @Summary      get all exercise activity
// @Description  get all exercise activity
// @Tags         activity exercise
// @Param        user_id path int true "User ID"
// @Param        date    query     string  true  "Date"
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
// @Tags         activity exercise
// @Param        user_id path int true "User ID"
// @Param        exercise_id path int true "Exercise Activity ID"
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
}
