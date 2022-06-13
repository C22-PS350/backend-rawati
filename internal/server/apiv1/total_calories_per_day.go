package apiv1

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/C22-PS350/backend-rawati/internal/utils"
	"github.com/go-chi/chi/v5"
)

var (
	selectTotalFoodCaloriesPerDay = `
		SELECT
			SUM(calories) as food_total
		FROM food_per_day
		JOIN users USING (user_id)
		WHERE user_id = ? AND food_date = ?
	`
	selectTotalExerciseCaloriesPerDay = `
		SELECT
			SUM(calories) as exercise_total
		FROM exercise_per_day
		JOIN users USING (user_id)
		WHERE user_id = ? AND exercise_date = ?
	`
)

// @Summary      get total calories per day
// @Description  get total calories per day
// @Tags         activity
// @Param        user_id  path  int                             true  "User ID"
// @Param        date     query  string  false  "Date (ISO 8601 - date only)"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=models.TotalCaloriesPerDayResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      403  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /users/{user_id}/total-calories [get]
func (h *Handler) GetTotalCaloriesPerDay(w http.ResponseWriter, r *http.Request) {
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
		y, m, d := time.Now().Date()
		date = fmt.Sprintf("%v-%v-%v", y, int(m), d)
	} else {
		_, err = time.Parse("2006-01-02", date)
		if err != nil {
			utils.RespondErr(w, http.StatusBadRequest, errors.New("invalid date query parameter value"))
			return
		}
	}

	var resp models.TotalCaloriesPerDayResponse
	if err := h.DB.Raw(selectTotalExerciseCaloriesPerDay, userID, date).Scan(&resp).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	if err := h.DB.Raw(selectTotalFoodCaloriesPerDay, userID, date).Scan(&resp).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	resp.UserID = userID
	utils.RespondOK(w, &resp)
}
