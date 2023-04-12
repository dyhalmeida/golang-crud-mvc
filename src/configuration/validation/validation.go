package validation

import (
	"encoding/json"
	"errors"

	restErrors "github.com/dyhalmeida/golang-crud-mvc/src/configuration/restErrors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
)

var transl ut.Translator

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {

		en := en.New()
		universalTranslator := ut.New(en, en)

		transl, _ = universalTranslator.GetTranslator("en")
		enTranslation.RegisterDefaultTranslations(val, transl)

	}
}

func ValidateError(err error) *restErrors.Error {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(err, &jsonErr) {

		return restErrors.NewBadRequestError("Invalid field type")

	} else if errors.As(err, &jsonValidationError) {

		errorsCauses := []restErrors.Causes{}

		for _, e := range err.(validator.ValidationErrors) {
			cause := restErrors.Causes{
				Message: e.Translate(transl),
				Field: e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return restErrors.NewBadRequestValidationError("Some fields are invalid", errorsCauses)

	} else {
		return restErrors.NewBadRequestError("Error trying to convert fields")
	}

}