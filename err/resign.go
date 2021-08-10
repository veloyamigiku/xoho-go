package err

import (
	"fmt"
	"xoho-go/model/json/enum"
)

type ResignError struct {
	Message string
	Code    enum.ResignCode
}

func (e *ResignError) Error() string {
	return e.Message
}

func (e *ResignError) Unwrap() error {
	return fmt.Errorf(e.Message)
}
