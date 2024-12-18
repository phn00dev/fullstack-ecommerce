package validate

import "github.com/go-playground/validator/v10"

func ValidateStruct(request any) error {
	validatorData := validator.New()
	return validatorData.Struct(request)
}
