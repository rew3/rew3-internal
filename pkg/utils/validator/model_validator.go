package validator

import (
	"fmt"

	"github.com/go-playground/validator"
)

/**
 * Model validator.
 */
type ModelValidator struct {
	validator *validator.Validate
}

func (mv *ModelValidator) ValidateModel(model interface{}) error {
	if err := mv.validator.Struct(model); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return fmt.Errorf("invalid validation error: %w", err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return fmt.Errorf("unexpected validation error type: %T", err)
		}
		errMsg := ""
		for _, validationErr := range errs {
			fieldName := validationErr.Field()
			errTag := validationErr.Tag()
			errMsg += fmt.Sprintf("Field: %s, Error: %s\n", fieldName, errTag)
		}
		return fmt.Errorf("validation error(s):\n%s", errMsg)
	}
	return nil
}
