package validation

type (
	ValidationError struct {
		Code string
	}
)

func (e *ValidationError) Error() string {
	return e.Code
}
