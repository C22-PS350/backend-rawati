package apiv1

import "testing"

var h *Handler

func TestMain(m *testing.M) {
	h = &Handler{
		DB: nil,
	}
	m.Run()
}
