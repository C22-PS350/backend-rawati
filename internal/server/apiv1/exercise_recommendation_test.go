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

func TestGetExerciseRecommendation(t *testing.T) {
	err := h.prepTestGetExerciseRecommendation()
	if err != nil {
		t.Fatalf("test getExerciseRecommendation prep error: %s", err)
	}

	p1 := models.ExerciseRecommendationRequest{
		Calories: 120,
	}

	b1 := bytes.Buffer{}
	if err := json.NewEncoder(&b1).Encode(&p1); err != nil {
		t.Fatalf("test getExerciseRecommendation prep error: %s", err)
	}

	p2 := models.ExerciseRecommendationRequest{
		Calories: 310,
	}

	b2 := bytes.Buffer{}
	if err := json.NewEncoder(&b2).Encode(&p2); err != nil {
		t.Fatalf("test getExerciseRecommendation prep error: %s", err)
	}

	endpoint := "/api/v1/recommendation/exercise"
	tests := []struct {
		Name            string
		Req             *http.Request
		Res             *httptest.ResponseRecorder
		ExpCode         int
		ExpExerciseName string
	}{
		{
			Name:    "Success",
			Req:     httptest.NewRequest(http.MethodPost, endpoint, &b1),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusOK,
			// ExpExerciseName: "basket",
		},
		{
			Name:    "Success",
			Req:     httptest.NewRequest(http.MethodPost, endpoint, &b2),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusOK,
			// ExpExerciseName: "badminton",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			h.CreateExerciseRecommendation(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
			}
			// var res models.ExerciseResponse
			// json.Unmarshal(test.Res.Body.Bytes(), &utils.JsonOK{Data: res})
			// if res.Name != test.ExpExerciseName {
			// 	t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
			// 	t.Errorf("expected exercise is %v, but got %v\n", test.ExpExerciseName, res.Name)
			// }
		})
	}
}

func (h *Handler) prepTestGetExerciseRecommendation() error {
	var e1, e2, e3 models.ExerciseTest1

	e1.Name = "badminton"
	e1.Calories = 280
	if err := faker.FakeData(&e1); err != nil {
		return err
	}

	e2.Name = "basket"
	e2.Calories = 120
	if err := faker.FakeData(&e2); err != nil {
		return err
	}

	e3.Name = "bal-balan"
	e3.Calories = 340
	if err := faker.FakeData(&e3); err != nil {
		return err
	}

	data := []models.ExerciseTest1{e1, e2, e3}
	if err := h.DB.Table("exercises").Create(&data).Error; err != nil {
		return err
	}

	return nil
}
