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
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Errors  []validationErr `json:"errors"`
}

type validationErr struct {
	Field   string `json:"field"`
	Message string `json:"message"`
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
		Errors:  []validationErr{},
	}

	errs := TranslateError(err)
	if len(errs) != 0 {
		e.Message = "request body validation error"
		e.Errors = append(e.Errors, errs...)
	}

	if code == http.StatusInternalServerError {
		fmt.Fprintf(os.Stderr, "internal server error: %s\n", err)
		e.Message = "server error: something went wrong"
	}

	d := JsonErr{
		Error: e,
	}
	writeToResponse(w, code, d)
}

func writeToResponse(w http.ResponseWriter, code int, body interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		panic(fmt.Sprintf("error encode json response: %s", err))
	}
}
