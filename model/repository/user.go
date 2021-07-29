package repository

import (
	"xoho-go/database"
	"xoho-go/model/db"
)

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
