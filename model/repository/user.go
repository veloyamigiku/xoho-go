package repository

import (
	"xoho-go/model/db"

	"gorm.io/gorm"
)

func UpdatePassword(
	db *gorm.DB,
	user *db.User,
	password string,
) error {

	result := db.
		Debug().
		Model(user).
		Update("password", password)

	return result.Error
}

func UpdateAuthMissCount(
	db *gorm.DB,
	userExt *db.UserExt,
	authMissCount int) error {

	result := db.
		Debug().
		Model(userExt).
		Where("Id = ?", userExt.Id).
		Update("auth_miss_count", authMissCount)
	return result.Error

}

func FindUserWithId(
	db *gorm.DB,
	id int) (
	user db.User,
	err error) {

	result := db.
		Debug().
		Joins("UserExt").
		First(&user, id)
	return user, result.Error

}

func FindUserWithName(
	db *gorm.DB,
	name string) (user db.User, err error) {

	result := db.
		Debug().
		Where("name = ?", name).
		Joins("UserExt").
		First(&user)
	return user, result.Error

}

func DeleteUserExt(
	db *gorm.DB,
	userExt *db.UserExt) error {

	result := db.
		Debug().
		Where("Id = ?", userExt.Id).
		Delete(userExt)

	return result.Error
}

func DeleteUser(
	db *gorm.DB,
	user *db.User) (int64, error) {

	result := db.
		Debug().
		Where("Id = ?", user.Id).
		Delete(user)
	return result.RowsAffected, result.Error
}

func ExistsUser(
	tx *gorm.DB,
	user *db.User) (bool, error) {

	var users []db.User
	exists := false
	result := tx.
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

func AddUser(
	db *gorm.DB,
	user *db.User) error {
	tx := db.
		Debug().
		Create(user)
	return tx.Error
}

func AddUserExt(
	db *gorm.DB,
	userExt *db.UserExt) error {
	tx := db.
		Debug().
		Create(userExt)
	return tx.Error
}

func AddAssociation(
	db *gorm.DB,
	user *db.User,
	userExt *db.UserExt) error {
	return db.
		Debug().
		Model(user).
		Association("UserExt").
		Append(userExt)
}
