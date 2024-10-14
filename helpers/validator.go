package helpers

import "github.com/go-playground/validator/v10"

type ValidatorService struct {
	Validator *validator.Validate
}

func (validatorService *ValidatorService) ValidateData(data interface{}) error {

	return validatorService.Validator.Struct(data)

}
func NewValidatorService() *ValidatorService {
	return &ValidatorService{
		Validator: validator.New(),
	}

}
