package json

import "xoho-go/model/json/enum"

type ResignRes struct {
	Status bool            `json:"status"`
	Code   enum.ResignCode `json:"code"`
	Msg    string          `json:"msg"`
}
