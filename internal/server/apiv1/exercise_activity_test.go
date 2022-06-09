package apiv1

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/bxcodec/faker/v3"
	"github.com/go-chi/chi/v5"
)

// func TestGetAllExerciseActivity(t *testing.T) {
// 	userID, date, err := h.prepTestGetAllExerciseActivity()
// 	if err != nil {
// 		t.Fatalf("test getAllExerciseActivity prep error: %s", err)
// 	}

// 	y, m, d := date.Year(), int(date.Month()), date.Day()
// 	endpoint := fmt.Sprintf("/api/v1/users/{user_id}/exercises?%v-%v-%v", y, m, d)
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

// 			h.GetAllExerciseActivity(test.Res, test.Req)
// 			if test.Res.Code != test.ExpCode {
// 				t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
// 				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
// 			}
// 		})
// 	}
// }

// func (h *Handler) prepTestGetAllExerciseActivity() (uint64, *time.Time, error) {
// 	var u1 models.ExerciseActivityTest1
// 	if err := faker.FakeData(&u1); err != nil {
// 		return 0, nil, err
// 	}

// 	var e1, e2 models.ExerciseActivityTest2
// 	e1.UserID, e2.UserID = u1.UserID, u1.UserID
// 	now := time.Now()
// 	e1.ExerciseDate, e2.ExerciseDate = now, now
// 	if err := faker.FakeData(&e1); err != nil {
// 		return 0, nil, err
// 	}
// 	if err := faker.FakeData(&e2); err != nil {
// 		return 0, nil, err
// 	}

// 	if err := h.DB.Table("users").Create(&u1).Error; err != nil {
// 		return 0, nil, err
// 	}

// 	data := []models.ExerciseActivityTest2{e1, e2}
// 	if err := h.DB.Table("exercise_per_day").Create(&data).Error; err != nil {
// 		return 0, nil, err
// 	}

// 	return u1.UserID, &now, nil
// }

func TestGetExerciseActivity(t *testing.T) {
	exerciseID, userID, err := h.prepTestGetExerciseActivity()
	if err != nil {
		t.Fatalf("test getExerciseActivity prep error: %s", err)
	}

	endpoint := "/api/v1/users/{user_id}/exercises/{exercise_id}"
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
			rctx.URLParams.Add("exercise_id", strconv.Itoa(int(exerciseID)))
			test.Req = test.Req.WithContext(context.WithValue(test.Req.Context(), chi.RouteCtxKey, rctx))

			h.GetExerciseActivity(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
			}
		})
	}
}

func (h *Handler) prepTestGetExerciseActivity() (uint64, uint64, error) {
	var u1 models.ExerciseActivityTest1
	if err := faker.FakeData(&u1); err != nil {
		return 0, 0, err
	}

	var e1 models.ExerciseActivityTest2
	e1.UserID, e1.ExerciseDate = u1.UserID, time.Now()
	if err := faker.FakeData(&e1); err != nil {
		return 0, 0, err
	}

	if err := h.DB.Table("users").Create(&u1).Error; err != nil {
		return 0, 0, err
	}

	if err := h.DB.Table("exercise_per_day").Create(&e1).Error; err != nil {
		return 0, 0, err
	}

	return e1.ExerciseActivityID, u1.UserID, nil
}
