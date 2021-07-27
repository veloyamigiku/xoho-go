package service

import (
	"sort"
	"xoho-go/model/convert"
	"xoho-go/model/db"
	"xoho-go/model/json"
	"xoho-go/model/repository"
)

func GetAllTheaters() []json.TheaterRes {
	theaters := repository.GetAllTheater()
	sort.Sort(theaters)

	theaterRes := []json.TheaterRes{}
	theaterPrefecture := []json.TheaterPrefecture{}
	theater := []json.Theater{}
	//preArea := db.Area{}
	var preArea *db.Area
	preArea = nil
	//prePrefecture := db.Prefecture{}
	var prePrefecture *db.Prefecture
	prePrefecture = nil
	for _, t := range theaters {
		if preArea != nil && (*preArea).Id != t.Area.Id {
			theaterRes = append(
				theaterRes,
				json.TheaterRes{
					Header: json.TheaterHeader{
						Title:  preArea.Name,
						Sub:    preArea.Sub,
						Option: "",
					},
					Prefecture: theaterPrefecture,
				},
			)
			theaterPrefecture = []json.TheaterPrefecture{}
			theater = []json.Theater{}
		} else if prePrefecture != nil && prePrefecture.Id != t.Prefecture.Id {
			theaterPrefecture = append(
				theaterPrefecture,
				json.TheaterPrefecture{
					Name:    prePrefecture.Name,
					Sub:     prePrefecture.Sub,
					Theater: theater,
				},
			)
			theater = []json.Theater{}
		}
		theater = append(
			theater,
			convert.ConvertTheaterFromDbToJson(t))
		//tmpArea := db.Area{}
		tmpArea := t.Area
		preArea = &tmpArea
		tmpPrefecture := t.Prefecture
		prePrefecture = &tmpPrefecture
	}
	if len(theaterPrefecture) > 0 {
		if len(theater) > 0 {
			theaterPrefecture = append(
				theaterPrefecture,
				json.TheaterPrefecture{
					Name:    prePrefecture.Name,
					Sub:     prePrefecture.Sub,
					Theater: theater,
				},
			)
		}
		theaterRes = append(
			theaterRes,
			json.TheaterRes{
				Header: json.TheaterHeader{
					Title:  preArea.Name,
					Sub:    preArea.Sub,
					Option: "",
				},
				Prefecture: theaterPrefecture,
			},
		)
	}

	return theaterRes
}

func GetAllTypeTheaters() []json.TheaterRes {
	allTypeIds := []int{6, 1, 2, 3, 4}
	res := []json.TheaterRes{}
	for _, allTypeId := range allTypeIds {
		var theaterHeader *json.TheaterHeader = nil
		prefectureIdTheatersMap := make(map[int][]db.Theater)
		theaters := repository.GetTheaterWithTypeId(allTypeId)
		sort.Sort(theaters)
		for _, theater := range theaters {
			// TheaterHeaderを作成する。
			if theaterHeader == nil {
				for _, typ := range theater.Type {
					if typ.Id == allTypeId {
						theaterHeader = &json.TheaterHeader{
							Title:  typ.Title,
							Sub:    typ.Sub,
							Option: typ.Opt,
						}
						break
					}
				}
			}
			prefectureIdTheatersMap[theater.PrefectureId] = append(prefectureIdTheatersMap[theater.PrefectureId], theater)
		}

		// TheaterPrefectureを作成する。
		theaterPrefecture := []json.TheaterPrefecture{}
		prefectureIds := []int{}
		for k := range prefectureIdTheatersMap {
			prefectureIds = append(prefectureIds, k)
		}
		sort.Slice(
			prefectureIds,
			func(i, j int) bool {
				return prefectureIds[i] < prefectureIds[j]
			})
		for _, prefectureId := range prefectureIds {
			oneTheaterOfPrefecture := prefectureIdTheatersMap[prefectureId][0]
			jsonTheaters := convert.ConvertTheatersFromDbToJson(prefectureIdTheatersMap[prefectureId])
			theaterPrefecture = append(
				theaterPrefecture,
				json.TheaterPrefecture{
					Name:    oneTheaterOfPrefecture.Prefecture.Name,
					Sub:     oneTheaterOfPrefecture.Prefecture.Sub,
					Theater: jsonTheaters,
				},
			)
		}

		// TheaterResを作成する。
		theaterRes := json.TheaterRes{
			Header:     *theaterHeader,
			Prefecture: theaterPrefecture,
		}
		res = append(res, theaterRes)
	}
	return res
}
