package apiv1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/bxcodec/faker/v3"
)

func TestGetFoodRecommendation(t *testing.T) {
	err := h.prepTestGetFoodRecommendation()
	if err != nil {
		t.Fatalf("test getFoodRecommendation prep error: %s", err)
	}

	p1 := models.FoodRecommendationRequest{
		Calories: 120,
	}

	b1 := bytes.Buffer{}
	if err := json.NewEncoder(&b1).Encode(&p1); err != nil {
		t.Fatalf("test getFoodRecommendation prep error: %s", err)
	}

	p2 := models.FoodRecommendationRequest{
		Calories: 310,
	}

	b2 := bytes.Buffer{}
	if err := json.NewEncoder(&b2).Encode(&p2); err != nil {
		t.Fatalf("test getFoodRecommendation prep error: %s", err)
	}

	endpoint := "/api/v1/recommendation/food"
	tests := []struct {
		Name        string
		Req         *http.Request
		Res         *httptest.ResponseRecorder
		ExpCode     int
		ExpFoodName string
	}{
		{
			Name:    "Success",
			Req:     httptest.NewRequest(http.MethodPost, endpoint, &b1),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusOK,
			// ExpFoodName: "omelette",
		},
		{
			Name:    "Success",
			Req:     httptest.NewRequest(http.MethodPost, endpoint, &b2),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusOK,
			// ExpFoodName: "sate rembiga",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			h.CreateFoodRecommendation(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
			}
			// var res models.FoodResponse
			// json.Unmarshal(test.Res.Body.Bytes(), &utils.JsonOK{Data: res})
			// if res.Name != test.ExpFoodName {
			// 	t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
			// 	t.Errorf("expected exercise is %v, but got %v\n", test.ExpFoodName, res.Name)
			// }
		})
	}
}

func (h *Handler) prepTestGetFoodRecommendation() error {
	var f1, f2, f3 models.FoodTest1

	f1.Name = "ayam goreng"
	f1.Calories = 280
	if err := faker.FakeData(&f1); err != nil {
		return err
	}

	f2.Name = "omelette"
	f2.Calories = 120
	if err := faker.FakeData(&f2); err != nil {
		return err
	}

	f3.Name = "sate rembiga"
	f3.Calories = 340
	if err := faker.FakeData(&f3); err != nil {
		return err
	}

	data := []models.FoodTest1{f1, f2, f3}
	if err := h.DB.Table("foods").Create(&data).Error; err != nil {
		return err
	}

	return nil
}
