package apiv1

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllExercises(t *testing.T) {
	endpoint := "/api/v1/exercises"
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
			h.GetAllExercises(test.Res, test.Req)
			if test.Res.Code != test.ExpCode {
				t.Errorf("%s error response: %s", test.Name, test.Res.Body.String())
				t.Errorf("expected code is %v, but got %v\n", test.ExpCode, test.Res.Code)
			}
		})
	}
}
