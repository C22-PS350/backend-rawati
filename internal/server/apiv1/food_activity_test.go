package apiv1

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/bxcodec/faker/v3"
	"github.com/go-chi/chi/v5"
)

// func TestGetAllFoodActivity(t *testing.T) {
// 	userID, date, err := h.prepTestGetAllFoodActivity()
// 	if err != nil {
// 		t.Fatalf("test getAllFoodActivity prep error: %s", err)
// 	}

// 	y, m, d := date.Year(), int(date.Month()), date.Day()
// 	endpoint := fmt.Sprintf("/api/v1/users/{user_id}/foods?%v-%v-%v", y, m, d)
// 	tests := []struct {
// 		Name    string
// 		Req     *http.Request
// 		Res     *httptest.ResponseRecorder
// 		ExpCode int
// 	}{
// 		{
// 			Name:    "Success",
// 			Req:     httptest.NewRequest(http.MethodGet, endpoint, nil),
// 			Res:     httptest.NewRecorder(),
// 			ExpCode: http.StatusOK,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			rctx := chi.NewRouteContext()
// 			rctx.URLParams.Add("user_id", strconv.Itoa(int(userID)))
// 			test.Req = test.Req.WithContext(context.WithValue(test.Req.Context(), chi.RouteCtxKey, rctx))

// 			h.GetAllFoodActivity(test.Res, test.Req)
// 			if test.Res.Code != test.ExpCode {
// 				t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
// 				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
// 			}
// 		})
// 	}
// }

// func (h *Handler) prepTestGetAllFoodActivity() (uint64, *time.Time, error) {
// 	var u1 models.FoodActivityTest1
// 	if err := faker.FakeData(&u1); err != nil {
// 		return 0, nil, err
// 	}

// 	var f1, f2 models.FoodActivityTest2
// 	f1.UserID, f2.UserID = u1.UserID, u1.UserID
// 	now := time.Now()
// 	f1.FoodDate, f2.FoodDate = now, now
// 	if err := faker.FakeData(&f1); err != nil {
// 		return 0, nil, err
// 	}
// 	if err := faker.FakeData(&f2); err != nil {
// 		return 0, nil, err
// 	}

// 	if err := h.DB.Table("users").Create(&u1).Error; err != nil {
// 		return 0, nil, err
// 	}

// 	data := []models.FoodActivityTest2{f1, f2}
// 	if err := h.DB.Table("food_per_day").Create(&data).Error; err != nil {
// 		return 0, nil, err
// 	}

// 	return u1.UserID, &now, nil
// }

func TestGetFoodActivity(t *testing.T) {
	foodID, userID, err := h.prepTestGetFoodActivity()
	if err != nil {
		t.Fatalf("test getFoodActivity prep error: %s", err)
	}

	endpoint := "/api/v1/users/{user_id}/foods/{food_id}"
	tests := []struct {
		Name    string
		Req     *http.Request
		Res     *httptest.ResponseRecorder
		ExpCode int
	}{
		{
			Name:    "Success",
			Req:     httptest.NewRequest(http.MethodGet, endpoint, nil),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("user_id", strconv.Itoa(int(userID)))
			rctx.URLParams.Add("food_id", strconv.Itoa(int(foodID)))
			test.Req = test.Req.WithContext(context.WithValue(test.Req.Context(), chi.RouteCtxKey, rctx))

			h.GetFoodActivity(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
			}
		})
	}
}

func (h *Handler) prepTestGetFoodActivity() (uint64, uint64, error) {
	var u1 models.FoodActivityTest1
	if err := faker.FakeData(&u1); err != nil {
		return 0, 0, err
	}

	var f1 models.FoodActivityTest2
	f1.UserID, f1.FoodDate = u1.UserID, time.Now()
	if err := faker.FakeData(&f1); err != nil {
		return 0, 0, err
	}

	if err := h.DB.Table("users").Create(&u1).Error; err != nil {
		return 0, 0, err
	}

	if err := h.DB.Table("food_per_day").Create(&f1).Error; err != nil {
		return 0, 0, err
	}

	return f1.FoodActivityID, u1.UserID, nil
}

func TestCreateFoodActivity(t *testing.T) {
	userData, userID, err := h.prepTestCreateFoodActivity()
	if err != nil {
		t.Fatalf("test createFoodActivity prep error: %s", err)
	}

	p1 := models.FoodActivityRequest{
		Name:     userData.Name,
		Calories: userData.Calories,
	}

	b1 := bytes.Buffer{}
	if err := json.NewEncoder(&b1).Encode(&p1); err != nil {
		t.Fatalf("test createFoodActivity prep error: %s", err)
	}

	p2 := models.FoodActivityRequest{
		Name: userData.Name,
	}

	b2 := bytes.Buffer{}
	if err := json.NewEncoder(&b2).Encode(&p2); err != nil {
		t.Fatalf("test createFoodActivity prep error: %s", err)
	}

	endpoint := "/api/v1/users/{user_id}/foods"
	tests := []struct {
		Name    string
		Req     *http.Request
		Res     *httptest.ResponseRecorder
		ExpCode int
	}{
		{
			Name:    "Success",
			Req:     httptest.NewRequest(http.MethodPost, endpoint, &b1),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusOK,
		},
		{
			Name:    "BadRequest",
			Req:     httptest.NewRequest(http.MethodPost, endpoint, &b2),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("user_id", strconv.Itoa(int(userID)))
			test.Req = test.Req.WithContext(context.WithValue(test.Req.Context(), chi.RouteCtxKey, rctx))

			h.CreateFoodActivity(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
			}
		})
	}
}

func (h *Handler) prepTestCreateFoodActivity() (*models.FoodActivityTest2, uint64, error) {
	var u1 models.FoodActivityTest1
	if err := faker.FakeData(&u1); err != nil {
		return nil, 0, err
	}

	var f1 models.FoodActivityTest2
	if err := faker.FakeData(&f1); err != nil {
		return nil, 0, err
	}

	if err := h.DB.Table("users").Create(&u1).Error; err != nil {
		return nil, 0, err
	}

	return &f1, u1.UserID, nil
}
