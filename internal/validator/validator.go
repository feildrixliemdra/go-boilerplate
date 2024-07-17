package validator

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"go-boilerplate/internal/payload"

	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

var TranslatorInst ut.Translator

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		english := en.New()
		uni := ut.New(english, english)
		trans, _ := uni.GetTranslator("en")
		TranslatorInst = trans

		_ = en_translations.RegisterDefaultTranslations(v, TranslatorInst)

		// register tag e.Field() use json tag
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

func TranslateErrorValidator(err error) (res []payload.ErrorValidation) {
	var (
		invalidValidationError *validator.InvalidValidationError
		errs                   validator.ValidationErrors
	)

	// not accepted error by validator
	if errors.As(err, &invalidValidationError) {
		return
	}

	// iterate error message
	ok := errors.As(err, &errs)
	if !ok {
		return
	}

	for _, e := range errs {
		res = append(res, payload.ErrorValidation{
			Field:   e.Field(),
			Message: e.Translate(TranslatorInst),
		})

	}

	return
}
