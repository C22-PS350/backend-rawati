package apiv1

import (
	"net/http"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/C22-PS350/backend-rawati/internal/utils"
)

// @Summary      get all foods
// @Description  get all foods
// @Tags         resources
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=[]models.FoodResponse}
// @Failure      500  {object}  utils.JsonErr
// @Router       /foods [get]
func (h *Handler) GetAllFoods(w http.ResponseWriter, r *http.Request) {
	resp := make([]models.FoodResponse, 0)
	if err := h.DB.Table("foods").Find(&resp).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondOK(w, &resp)
}
