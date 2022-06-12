package utils

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	trans ut.Translator
)

func TranslateError(err error) (errs []validationErr) {
	validatorErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return []validationErr{}
	}

	var errInfo validationErr
	for _, e := range validatorErrs {
		errInfo.Field = e.Field()
		errInfo.Message = e.Translate(trans)
		errs = append(errs, errInfo)
	}

	return errs
}

func SetupTranslator(t ut.Translator) {
	trans = t
}
