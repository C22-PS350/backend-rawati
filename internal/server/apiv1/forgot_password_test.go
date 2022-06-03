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

func TestForgotPassword(t *testing.T) {
	email, err := h.prepTestForgotPassword()
	if err != nil {
		t.Fatalf("test forgot password prep error: %s", err)
	}

	p1 := models.ForgotPasswordRequest{
		Email: email,
	}

	b1 := bytes.Buffer{}
	if err := json.NewEncoder(&b1).Encode(&p1); err != nil {
		t.Fatalf("test forgot password prep error: %s", err)
	}

	p2 := models.ForgotPasswordRequest{
		Email: "somethingwrong@gmail.com",
	}

	b2 := bytes.Buffer{}
	if err := json.NewEncoder(&b2).Encode(&p2); err != nil {
		t.Fatalf("test forgot password prep error: %s", err)
	}

	endpoint := "/api/v1/auth/forgot-password"
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
			Name:    "UserNotFound",
			Req:     httptest.NewRequest(http.MethodPut, endpoint, &b2),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			h.ForgotPassword(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
			}
		})
	}
}

func (h *Handler) prepTestForgotPassword() (string, error) {
	var f1 models.ForgotPasswordTest1
	if err := faker.FakeData(&f1); err != nil {
		return "", err
	}

	if err := h.DB.Table("users").Create(&f1).Error; err != nil {
		return "", err
	}
	return f1.Email, nil
}
