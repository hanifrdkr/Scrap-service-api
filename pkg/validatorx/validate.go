package validatorx

import "github.com/go-playground/validator/v10"

func Validate(param interface{}) (err error) {
	v := validator.New()
	err = v.Struct(param)
	if err != nil {
		return
	}
	return
}
