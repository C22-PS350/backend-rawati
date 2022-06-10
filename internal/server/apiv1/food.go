package apiv1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/C22-PS350/backend-rawati/internal/utils"
)

// @Summary      get all foods
// @Description  get all foods
// @Tags         resources
// @Param        page   query  int  false  "Page"   minimum(1)
// @Param        limit  query  int  false  "Limit"  minimum(1)  maximum(20)
// @Produce      json
// @Success      200  {object}  utils.JsonOK{data=[]models.FoodResponse}
// @Failure      400  {object}  utils.JsonErr
// @Failure      500  {object}  utils.JsonErr
// @Router       /foods [get]
func (h *Handler) GetAllFoods(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "1"
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("invalid query string 'page' value"))
		return
	}

	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limitStr = "20"
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		utils.RespondErr(w, http.StatusBadRequest, errors.New("invalid query string 'limit' value"))
		return
	}

	if page <= 0 {
		page = 1
	}

	switch {
	case limit > 20:
		limit = 20
	case limit <= 0:
		limit = 20
	}

	resp := make([]models.FoodResponse, 0)
	if err := h.DB.Table("foods").Offset((page - 1) * limit).Limit(limit).Find(&resp).Error; err != nil {
		utils.RespondErr(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondOK(w, &resp)
}
