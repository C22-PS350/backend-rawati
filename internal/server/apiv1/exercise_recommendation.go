package apiv1

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net/http"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/C22-PS350/backend-rawati/internal/utils"
	"github.com/go-playground/validator/v10"
)

// @Summary      get exercise recommendation
// @Description  get exercise recommendation
// @Tags         recommendation
// @Accept       json
// @Param        payload  body  models.ExerciseRecommendationRequest  true  "request body"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=[]models.ExerciseResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /recommendation/exercise [post]
func (h *Handler) CreateExerciseRecommendation(w http.ResponseWriter, r *http.Request) {
	var req models.ExerciseRecommendationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error decoding request: %s", err))
		return
	}

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("request body validation error"))
		return
	}

	data := make([]models.ExerciseResponse, 0)
	if err := h.DB.Table("exercises").Find(&data).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	rand.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})

	resp := make([]models.ExerciseResponse, 0, 3)
	for _, v := range data {
		if math.Abs(math.Abs(req.Calories)-math.Abs(v.Calories)) <= 50 {
			if len(resp) == 3 {
				break
			}
			resp = append(resp, v)
		}
	}

	utils.RespondOK(w, &resp)
}
