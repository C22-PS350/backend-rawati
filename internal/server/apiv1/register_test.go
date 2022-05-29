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

func TestRegister(t *testing.T) {
	userData, err := h.PrepTestRegister()
	if err != nil {
		t.Fatalf("test login prep error: %s", err)
	}

	p1 := models.RegisterRequest{
		Name:     &userData.Name,
		Username: &userData.Username,
		Email:    &userData.Email,
		Password: &userData.Password,
	}

	b1 := bytes.Buffer{}
	if err := json.NewEncoder(&b1).Encode(&p1); err != nil {
		t.Fatalf("test login prep error: %s", err)
	}

	p2 := models.RegisterRequest{
		Name:     &userData.Name,
		Username: &userData.Username,
	}

	b2 := bytes.Buffer{}
	if err := json.NewEncoder(&b2).Encode(&p2); err != nil {
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
			Req:     httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", &b1),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusOK,
		},
		{
			Name:    "IncompleteInput",
			Req:     httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", &b2),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			h.Register(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
			}
		})
	}
}

func (h *Handler) PrepTestRegister() (*models.RegisterTest1, error) {
	var r1 models.RegisterTest1
	if err := faker.FakeData(&r1); err != nil {
		return nil, err
	}
	return &r1, nil
}
