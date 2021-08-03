package json

import (
	"xoho-go/model/json/enum"
)

type UpdatePasswordRes struct {
	Status bool                    `json:"status"`
	Code   enum.UpdatePasswordCode `json:"code"`
}
