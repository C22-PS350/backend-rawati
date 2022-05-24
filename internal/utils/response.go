package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

func RespondOK(w http.ResponseWriter, payload interface{}) {
	d := JsonOK{
		Data: payload,
	}
	writeToResponse(w, http.StatusOK, d)
}

func RespondErr(w http.ResponseWriter, code int, err error) {
	e := errContent{
		Code:    code,
		Message: err.Error(),
	}

	d := JsonErr{
		Error: e,
	}
	writeToResponse(w, code, d)

	if code == http.StatusInternalServerError {
		fmt.Fprintf(os.Stderr, "internal server error: %s\n", err)
	}
}

func writeToResponse(w http.ResponseWriter, code int, body interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		panic(fmt.Sprintf("error encode json response: %s", err))
	}
}
