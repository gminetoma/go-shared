package validation

import "net/mail"

type (
	Validator struct {
		errors []error
	}
)

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Errors() []error {
	return v.errors
}

func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}

func (v *Validator) AddError(code string) {
	v.errors = append(v.errors, &ValidationError{Code: code})
}

func (v *Validator) RequiredString(value, code string) {
	if value == "" {
		v.AddError(code)
	}
}

func (v *Validator) MinLength(value string, min int, code string) {
	if value != "" && len(value) < min {
		v.AddError(code)
	}
}

func (v *Validator) MaxLength(value string, max int, code string) {
	if value != "" && len(value) > max {
		v.AddError(code)
	}
}

func (v *Validator) ValidEmail(value, code string) {
	if value == "" {
		return
	}

	if _, err := mail.ParseAddress(value); err != nil {
		v.AddError(code)
	}
}
