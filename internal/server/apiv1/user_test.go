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
	"golang.org/x/crypto/bcrypt"
)

func TestGetUser(t *testing.T) {
	userID, err := h.prepTestGetUser()
	if err != nil {
		t.Fatalf("test getUser prep error: %s", err)
	}

	endpoint := "/api/v1/users/{user_id}"
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

			h.GetUser(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
			}
		})
	}
}

func (h *Handler) prepTestGetUser() (uint64, error) {
	var g1 models.UpdateUserTest1
	if err := faker.FakeData(&g1); err != nil {
		return 0, err
	}

	if err := h.DB.Table("users").Create(&g1).Error; err != nil {
		return 0, err
	}

	return g1.UserID, nil
}

func TestUpdateUser(t *testing.T) {
	userData, userID, err := h.prepTestUpdateUser()
	if err != nil {
		t.Fatalf("test updateUser prep error: %s", err)
	}

	p1 := models.UpdateUserRequest{
		Name:     &userData.Name,
		Username: &userData.Username,
		Email:    &userData.Email,
	}

	b1 := bytes.Buffer{}
	if err := json.NewEncoder(&b1).Encode(&p1); err != nil {
		t.Fatalf("test updateUser prep error: %s", err)
	}

	p2 := models.UpdateUserRequest{
		Username: &userData.Username,
		Email:    &userData.Email,
	}

	b2 := bytes.Buffer{}
	if err := json.NewEncoder(&b2).Encode(&p2); err != nil {
		t.Fatalf("test updateUser prep error: %s", err)
	}

	endpoint := "/api/v1/users/{user_id}/update"
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
		{
			Name:    "IncompleteInput",
			Req:     httptest.NewRequest(http.MethodPut, endpoint, &b2),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("user_id", strconv.Itoa(int(userID)))
			test.Req = test.Req.WithContext(context.WithValue(test.Req.Context(), chi.RouteCtxKey, rctx))

			h.UpdateUser(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
			}
		})
	}
}

func (h *Handler) prepTestUpdateUser() (*models.UpdateUserTest1, uint64, error) {
	var u1 models.UpdateUserTest1
	if err := faker.FakeData(&u1); err != nil {
		return nil, 0, err
	}

	var u2 models.UpdateUserTest1
	if err := faker.FakeData(&u2); err != nil {
		return nil, 0, err
	}

	if err := h.DB.Table("users").Create(&u1).Error; err != nil {
		return nil, 0, err
	}

	return &u2, u1.UserID, nil
}

func TestUpdateUserPassword(t *testing.T) {
	userID, plainPwd, err := h.prepUpdateUserPassword()
	if err != nil {
		t.Fatalf("test updateUserPassword prep error: %s", err)
	}

	p1 := models.UpdateUserPwdRequest{
		OldPassword: plainPwd,
		NewPassword: "newPassword",
	}

	b1 := bytes.Buffer{}
	if err := json.NewEncoder(&b1).Encode(&p1); err != nil {
		t.Fatalf("test updateUserPassword prep error: %s", err)
	}

	p2 := models.UpdateUserPwdRequest{
		OldPassword: "somethingwrong",
		NewPassword: "newPassword",
	}

	b2 := bytes.Buffer{}
	if err := json.NewEncoder(&b2).Encode(&p2); err != nil {
		t.Fatalf("test updateUserPassword prep error: %s", err)
	}

	endpoint := "/api/v1/users/{user_id}/update-password"
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
		{
			Name:    "WrongOldPassword",
			Req:     httptest.NewRequest(http.MethodPut, endpoint, &b2),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("user_id", strconv.Itoa(int(userID)))
			test.Req = test.Req.WithContext(context.WithValue(test.Req.Context(), chi.RouteCtxKey, rctx))

			h.UpdateUserPassword(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
			}
		})
	}
}

func (h *Handler) prepUpdateUserPassword() (uint64, string, error) {
	var u1 models.UpdateUserTest1
	if err := faker.FakeData(&u1); err != nil {
		return 0, "", err
	}

	plainPwd := u1.Password
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(u1.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, "", err
	}

	u1.Password = string(hashedPwd)
	if err := h.DB.Table("users").Create(&u1).Error; err != nil {
		return 0, "", err
	}

	return u1.UserID, plainPwd, nil
}

func TestDeleteUser(t *testing.T) {
	userID, err := h.prepTestDeleteUser()
	if err != nil {
		t.Fatalf("test deleteUser prep error: %s", err)
	}

	endpoint := "/api/v1/users/{user_id}"
	tests := []struct {
		Name    string
		Req     *http.Request
		Res     *httptest.ResponseRecorder
		ExpCode int
	}{
		{
			Name:    "Success",
			Req:     httptest.NewRequest(http.MethodDelete, endpoint, nil),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("user_id", strconv.Itoa(int(userID)))
			test.Req = test.Req.WithContext(context.WithValue(test.Req.Context(), chi.RouteCtxKey, rctx))

			h.DeleteUser(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
			}
		})
	}

}

func (h *Handler) prepTestDeleteUser() (uint64, error) {
	var d1 models.UpdateUserTest1
	if err := faker.FakeData(&d1); err != nil {
		return 0, err
	}

	if err := h.DB.Table("users").Create(&d1).Error; err != nil {
		return 0, err
	}

	return d1.UserID, nil
}
