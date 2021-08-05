package err

import (
	"fmt"
	"xoho-go/model/json/enum"
)

type UpdatePasswordError struct {
	Message string
	Code    enum.UpdatePasswordCode
}

func (e *UpdatePasswordError) Error() string {
	return e.Message
}

func (e *UpdatePasswordError) Unwrap() error {
	return fmt.Errorf(e.Message)
}
