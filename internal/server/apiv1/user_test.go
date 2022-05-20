package apiv1

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	p := strings.NewReader(`{"name": "farryl", "password": "farryl"}`)

	tests := []struct {
		Name    string
		Req     *http.Request
		Res     *httptest.ResponseRecorder
		ExpCode int
	}{
		{
			Name:    "Success",
			Req:     httptest.NewRequest(http.MethodPost, "/api/v1/users", p),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusOK,
		},
		{
			Name:    "Internal Server Error",
			Req:     httptest.NewRequest(http.MethodPost, "/api/v1/users", nil),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			h.CreateUser(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Fail()
			}
		})
	}
}
