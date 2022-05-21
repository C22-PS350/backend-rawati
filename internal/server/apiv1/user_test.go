package apiv1

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	p1 := strings.NewReader(`{"name": "farryl", "password": "azmi"}`)
	p2 := strings.NewReader(`{"name": "farryl"}`)

	tests := []struct {
		Name    string
		Req     *http.Request
		Res     *httptest.ResponseRecorder
		ExpCode int
	}{
		{
			Name:    "Success",
			Req:     httptest.NewRequest(http.MethodPost, "/api/v1/users", p1),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusOK,
		},
		{
			Name:    "Internal Server Error",
			Req:     httptest.NewRequest(http.MethodPost, "/api/v1/users", nil),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusInternalServerError,
		},
		{
			Name:    "Bad Request",
			Req:     httptest.NewRequest(http.MethodPost, "/api/v1/users", p2),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			h.CreateUser(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
			}
		})
	}
}
