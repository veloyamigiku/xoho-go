package convert

import (
	"xoho-go/model/db"
	"xoho-go/model/json"
)

func ConvertTheaterFromDbToJson(theater db.Theater) json.Theater {
	res := json.Theater{}
	res.Name = theater.Name
	res.Sub = theater.Sub
	res.Url = theater.Url
	res.Type = ConvertTypeListFromDbToJson(theater.Type)
	return res
}

func ConvertTheatersFromDbToJson(theaters []db.Theater) []json.Theater {
	res := []json.Theater{}
	for _, theater := range theaters {
		jsonTheater := ConvertTheaterFromDbToJson(theater)
		res = append(res, jsonTheater)
	}
	return res
}
