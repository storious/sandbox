package utils

import (
	"errors"
	"fmt"
	"sandbox/config"

	"github.com/go-playground/validator/v10"
)

func Validate(cfg *config.Config) bool {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(cfg)
	if err != nil {

		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			fmt.Println(err)
			return false
		}

		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			for _, e := range validateErrs {
				fmt.Println(e.Namespace())
				fmt.Println(e.Field())
				fmt.Println(e.StructNamespace())
				fmt.Println(e.StructField())
				fmt.Println(e.Tag())
				fmt.Println(e.ActualTag())
				fmt.Println(e.Kind())
				fmt.Println(e.Type())
				fmt.Println(e.Value())
				fmt.Println(e.Param())
				fmt.Println()
			}
			return false
		}
	}
	return !check(cfg.MaxCPUTime) || !check(cfg.MaxRealTime) || !check(cfg.MaxMemory) // for unlimited cpu or real tim ,memory
}

func check[T int | int64](val T) bool {
	return val >= 1 || val == config.UNLIMITED
}
