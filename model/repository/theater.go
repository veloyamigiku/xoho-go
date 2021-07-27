package repository

import (
	"xoho-go/database"
	"xoho-go/model/db"
)

func GetTheaterWithTypeId(typeId int) db.TheaterByAreaPref {
	theaters := db.TheaterByAreaPref{}
	database.
		DB.
		Debug().
		Model(&theaters).
		Joins("join theater_types on theaters.id = theater_types.theater_id and theater_types.type_id = ?", typeId).
		Joins("join types on theater_types.type_id = types.id").
		Preload("Type").
		Joins("Prefecture").
		Joins("Area").
		Find(&theaters)
	return theaters
}

func GetAllTheater() db.TheaterByAreaPref {
	theaters := db.TheaterByAreaPref{}
	database.
		DB.
		Debug().
		Model(&theaters).
		Preload("Type").
		Joins("Prefecture").
		Joins("Area").
		Find(&theaters)
	return theaters
}
