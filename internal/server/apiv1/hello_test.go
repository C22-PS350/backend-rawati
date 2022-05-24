package apiv1

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		Name    string
		Req     *http.Request
		Res     *httptest.ResponseRecorder
		ExpCode int
	}{
		{
			Name:    "Success",
			Req:     httptest.NewRequest(http.MethodGet, "/api/v1/hello", nil),
			Res:     httptest.NewRecorder(),
			ExpCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			h.Hello(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
			}
		})
	}
}
