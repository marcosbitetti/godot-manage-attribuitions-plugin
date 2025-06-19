package usecases

// ErrMissingArgument represents an error indicating that a required argument is missing.
type ErrMissingArgument struct{}

func (e ErrMissingArgument) Error() string {
	return "missing argument"
}

func NewErrMissingArgument() error {
	return ErrMissingArgument{}
}

// ErrInvalidValue represents an error indicating that an argument has an invalid value.
type ErrInvalidValue struct{}

func (e ErrInvalidValue) Error() string {
	return "invalid value"
}

func NewErrInvalidValue() error {
	return ErrInvalidValue{}
}

