package apiv1

import (
	"net/http"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/C22-PS350/backend-rawati/internal/utils"
)

// @Summary      get all exercises
// @Description  get all exercises
// @Tags         resources
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=[]models.ExerciseResponse}
// @Failure      500  {object}  utils.JsonErr
// @Router       /exercises [get]
func (h *Handler) GetAllExercises(w http.ResponseWriter, r *http.Request) {
	resp := make([]models.ExerciseResponse, 0)
	if err := h.DB.Table("exercises").Find(&resp).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondOK(w, &resp)
}
