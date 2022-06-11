package apiv1

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/bxcodec/faker/v3"
	"github.com/go-chi/chi/v5"
)

func TestGetUserProfile(t *testing.T) {
	userID, err := h.prepTestGetUserProfile()
	if err != nil {
		t.Fatalf("test getUserProfile prep error: %s", err)
	}

	endpoint := "/api/v1/users/{user_id}/profile"
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
			test.Req = test.Req.WithContext(context.WithValue(test.Req.Context(), chi.RouteCtxKey, rctx))

			h.GetUserProfile(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
			}
		})
	}
}

func (h *Handler) prepTestGetUserProfile() (uint64, error) {
	var u1 models.UserProfileTest1
	if err := faker.FakeData(&u1); err != nil {
		return 0, err
	}

	var u2 models.UserProfileTest2
	if err := faker.FakeData(&u2); err != nil {
		return 0, err
	}

	if err := h.DB.Table("users").Create(&u1).Error; err != nil {
		return 0, err
	}

	u2.UserID = u1.UserID
	if err := h.DB.Table("user_profile").Create(&u2).Error; err != nil {
		return 0, err
	}

	return u1.UserID, nil
}

func TestUpdateUserProfile(t *testing.T) {
	userData, userID, err := h.prepTestUpdateUserProfile()
	if err != nil {
		t.Fatalf("test updateUserProfile prep error: %s", err)
	}

	p1 := models.UserProfileRequest{
		Height:     &userData.Height,
		Weight:     &userData.Weight,
		WeightGoal: &userData.WeightGoal,
		Gender:     &userData.Gender,
		BirthDate:  &userData.BirthDate,
	}

	b1 := bytes.Buffer{}
	if err := json.NewEncoder(&b1).Encode(&p1); err != nil {
		t.Fatalf("test updateUserProfile prep error: %s", err)
	}

	endpoint := "/api/v1/users/{user_id}/profile"
	tests := []struct {
		Name    string
		Req     *http.Request
		Res     *httptest.ResponseRecorder
		ExpCode int
	}{
		{
			Name:    "Success",
			Req:     httptest.NewRequest(http.MethodPut, endpoint, &b1),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("user_id", strconv.Itoa(int(userID)))
			test.Req = test.Req.WithContext(context.WithValue(test.Req.Context(), chi.RouteCtxKey, rctx))

			h.UpdateUserProfile(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
			}
		})
	}
}

func (h *Handler) prepTestUpdateUserProfile() (*models.UserProfileTest2, uint64, error) {
	var u1 models.UserProfileTest1
	if err := faker.FakeData(&u1); err != nil {
		return nil, 0, err
	}

	var u2 models.UserProfileTest2
	if err := faker.FakeData(&u2); err != nil {
		return nil, 0, err
	}

	var u3 models.UserProfileTest2
	if err := faker.FakeData(&u3); err != nil {
		return nil, 0, err
	}

	if err := h.DB.Table("users").Create(&u1).Error; err != nil {
		return nil, 0, err
	}

	u2.UserID = u1.UserID
	if err := h.DB.Table("user_profile").Create(&u2).Error; err != nil {
		return nil, 0, err
	}

	return &u3, u2.UserID, nil
}

// func TestCreateUserProfile(t *testing.T) {
// 	userData, err := h.prepTestCreateUserProfile()
// 	if err != nil {
// 		t.Fatalf("test create user profile prep error: %s", err)
// 	}

// 	p1 := models.UserProfileRequest{
// 		UserID:     userData.UserID,
// 		Height:     &userData.Height,
// 		Weight:     &userData.Weight,
// 		WeightGoal: &userData.WeightGoal,
// 		Gender:     &userData.Gender,
// 		BirthDate:  &userData.BirthDate,
// 	}

// 	b1 := bytes.Buffer{}
// 	if err := json.NewEncoder(&b1).Encode(&p1); err != nil {
// 		t.Fatalf("test create user profile prep error: %s", err)
// 	}

// 	endpoint := "/api/v1/users/profile"
// 	tests := []struct {
// 		Name    string
// 		Req     *http.Request
// 		Res     *httptest.ResponseRecorder
// 		ExpCode int
// 	}{
// 		{
// 			Name:    "Success",
// 			Req:     httptest.NewRequest(http.MethodPost, endpoint, &b1),
// 			Res:     httptest.NewRecorder(),
// 			ExpCode: http.StatusOK,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			h.CreateUserProfile(test.Res, test.Req)
// 			if test.Res.Code != test.ExpCode {
// 				t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
// 				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
// 			}
// 		})
// 	}
// }

// func (h *Handler) prepTestCreateUserProfile() (*models.UserProfileTest2, error) {
// 	var u1 models.UserProfileTest1
// 	if err := faker.FakeData(&u1); err != nil {
// 		return nil, err
// 	}

// 	var u2 models.UserProfileTest2
// 	if err := faker.FakeData(&u2); err != nil {
// 		return nil, err
// 	}

// 	if err := h.DB.Table("users").Create(&u1).Error; err != nil {
// 		return nil, err
// 	}

// 	u2.UserID = u1.UserID
// 	return &u2, nil
// }
