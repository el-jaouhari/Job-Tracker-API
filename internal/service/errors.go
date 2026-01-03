package service

import "errors"

var (
	ErrInvalidStatus = errors.New("invalid application status")
	ErrJobNotFound   = errors.New("job not found")
	ErrMissingField  = errors.New("missing required field")
	ErrInvalidID     = errors.New("invalid job ID")
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

type StatusError struct {
	Status      string
	ValidStatus []string
}

func (e *StatusError) Error() string {
	return "invalid application status: " + e.Status
}
