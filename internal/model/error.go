package model

import "fmt"

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("%s-%s-%s", e.FailedField, e.Tag, e.Value)
}
