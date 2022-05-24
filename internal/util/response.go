package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonOK struct {
	Data interface{} `json:"data"`
}

type JsonErr struct {
	Error errContent `json:"error"`
}

type errContent struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}

func WriteToResponse(w http.ResponseWriter, code int, body interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		panic(fmt.Sprintf("error encode json response: %s", err))
	}
}

func RespondOK(w http.ResponseWriter, payload interface{}) {
	d := JsonOK{
		Data: payload,
	}
	WriteToResponse(w, http.StatusOK, d)
}

func RespondErr(w http.ResponseWriter, code int, err error) {
	e := errContent{
		Code:    code,
		Message: err.Error(),
	}

	d := JsonErr{
		Error: e,
	}
	WriteToResponse(w, code, d)

	if code == http.StatusInternalServerError {
		fmt.Fprintf(w, "internal server error: %s\n", err)
	}
}
