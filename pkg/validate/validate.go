package validate

import "github.com/go-playground/validator/v10"

// validate holds the settings and caches for validating request struct values.
var validate *validator.Validate

func init() {

	// Instantiate a validator.
	validate = validator.New()
}

/*Validator struct is for storing the custom validator that will be registered to echo server */
type Validator struct { }

/*
Validate is struct method that is called by registered validator in echo to validate
*/
func (*Validator) Validate(i interface{}) error {
	return validate.Struct(i)
}


// Check validates the provided model against it's declared tags.
func Check(val interface{}) error {
	return  validate.Struct(val)
}
