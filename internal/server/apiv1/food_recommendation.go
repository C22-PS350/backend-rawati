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
)

// @Summary      get food recommendation
// @Description  get food recommendation
// @Tags         recommendation
// @Accept       json
// @Param        payload  body  models.FoodRecommendationRequest  true  "request body"
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=[]models.FoodResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /recommendation/food [post]
func (h *Handler) CreateFoodRecommendation(w http.ResponseWriter, r *http.Request) {
	var req models.FoodRecommendationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, fmt.Errorf("error decoding request: %s", err))
		return
	}

	if err := h.V.Struct(&req); err != nil {
		utils.RespondErr(w, http.StatusBadRequest, err)
		return
	}

	data := make([]models.FoodResponse, 0)
	if err := h.DB.Table("foods").Find(&data).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	rand.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})

	resp := make([]models.FoodResponse, 0, 3)
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
