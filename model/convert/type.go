package convert

import (
	"xoho-go/model/db"
	"xoho-go/model/json"
)

func ConvertTypeFromDbToJson(typ db.Type) json.Type {
	res := json.Type{}
	res.Name = typ.Name
	res.Sub = typ.Sub
	res.Title = typ.Title
	res.Option = typ.Opt
	res.Icon = []string{typ.IconPrefix, typ.IconClass}
	return res
}

func ConvertTypeListFromDbToJson(typList []db.Type) []json.Type {
	res := []json.Type{}
	for _, typ := range typList {
		jsonTyp := ConvertTypeFromDbToJson(typ)
		res = append(res, jsonTyp)
	}
	return res
}
