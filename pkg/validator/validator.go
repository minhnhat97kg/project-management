package validator

import (
	"project-management/internal/model"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(v interface{}) []*model.ErrorResponse {
	var errors []*model.ErrorResponse
	err := validate.Struct(v)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element model.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
