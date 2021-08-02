package repository

import (
	"xoho-go/database"
	"xoho-go/model/db"
)

func FindUserWithName(name string) (user db.User, err error) {

	result := database.
		DB.
		Debug().
		Where("name = ?", name).
		Joins("UserExt").
		First(&user)
	return user, result.Error

}

func ExistsUser(user *db.User) (bool, error) {

	var users []db.User
	exists := false
	result := database.
		DB.
		Debug().
		Where(
			"name = ?",
			user.Name).
		Find(&users)
	if result.Error == nil {
		if result.RowsAffected > 0 {
			exists = true
		}
	}
	return exists, result.Error

}

func AddUser(user *db.User) error {
	tx := database.
		DB.
		Debug().
		Create(user)
	return tx.Error
}

func AddUserExt(userExt *db.UserExt) error {
	tx := database.
		DB.
		Debug().
		Create(userExt)
	return tx.Error
}

func AddAssociation(user *db.User, userExt *db.UserExt) error {
	return database.
		DB.
		Debug().
		Model(user).
		Association("UserExt").
		Append(userExt)
}
