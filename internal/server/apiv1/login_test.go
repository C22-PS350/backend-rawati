package apiv1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/C22-PS350/backend-rawati/internal/models"
	"github.com/bxcodec/faker/v3"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	userData, plainPwd, err := h.PrepTestLogin()
	if err != nil {
		t.Fatalf("test login prep error: %s", err)
	}

	p1 := models.LoginRequest{
		Identifier: userData.Username,
		Password:   plainPwd,
	}

	b1 := bytes.Buffer{}
	if err := json.NewEncoder(&b1).Encode(&p1); err != nil {
		t.Fatalf("test login prep error: %s", err)
	}

	p2 := models.LoginRequest{
		Identifier: "somethingwrong",
		Password:   plainPwd,
	}

	b2 := bytes.Buffer{}
	if err := json.NewEncoder(&b2).Encode(&p2); err != nil {
		t.Fatalf("test login prep error: %s", err)
	}

	p3 := models.LoginRequest{
		Identifier: userData.Username,
		Password:   "somethingwrong",
	}

	b3 := bytes.Buffer{}
	if err := json.NewEncoder(&b3).Encode(&p3); err != nil {
		t.Fatalf("test login prep error: %s", err)
	}

	tests := []struct {
		Name    string
		Req     *http.Request
		Res     *httptest.ResponseRecorder
		ExpCode int
	}{
		{
			Name:    "Success",
			Req:     httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", &b1),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusOK,
		},
		{
			Name:    "WrongUsername",
			Req:     httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", &b2),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusBadRequest,
		},
		{
			Name:    "WrongPassword",
			Req:     httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", &b3),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			h.Login(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
			}
		})
	}

}

func (h *Handler) PrepTestLogin() (*models.LoginTest1, string, error) {
	var l1 models.LoginTest1
	var l2 models.LoginTest2
	if err := faker.FakeData(&l1); err != nil {
		return nil, "", err
	}

	l2.UserID = l1.UserID
	l2.CreatedAt = time.Now()
	if err := faker.FakeData(&l2); err != nil {
		return nil, "", err
	}

	plainPwd := l1.Password
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(plainPwd), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	l1.Password = string(hashedPwd)
	if err := h.DB.Table("users").Create(&l1).Error; err != nil {
		return nil, "", err
	}

	if err := h.DB.Table("user_token").Create(&l2).Error; err != nil {
		return nil, "", err
	}

	return &l1, plainPwd, nil
}
