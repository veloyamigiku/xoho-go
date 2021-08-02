package service

import (
	"crypto/sha256"
	"fmt"
	"xoho-go/database"
	"xoho-go/model/db"
	"xoho-go/model/json"
	"xoho-go/model/repository"
	"xoho-go/utils"

	"gorm.io/gorm"
)

func Login(login json.Login) error {

	err := database.DB.Transaction(func(tx *gorm.DB) error {

		var err error
		err = nil

		user, findUserErr := repository.FindUserWithName(login.Name)
		if findUserErr != nil {
			err = fmt.Errorf("error: find user")
			return err
		}

		hash := utils.HashPassword(login.Password)
		trueHash := user.Password
		if hash != trueHash {
			err = fmt.Errorf("error: not equal password")
			return err
		}

		return err
	})
	return err
}

func SignUp(signup json.Signup) error {

	err := database.DB.Transaction(func(tx *gorm.DB) error {

		userExt := &db.UserExt{
			AuthMissCount: 0,
		}
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(signup.Password)))

		user := &db.User{
			Name:     signup.Name,
			Password: hash,
			UserExt:  *userExt,
		}

		var err error
		err = nil

		exists, existsUserErr := repository.ExistsUser(user)
		if existsUserErr != nil || exists {
			err = fmt.Errorf("error: exists user")
			return err
		}

		addUserErr := repository.AddUser(user)
		if addUserErr != nil {
			err = fmt.Errorf("error: add user")
			return err
		}

		addUserExtErr := repository.AddUserExt(userExt)
		if addUserExtErr != nil {
			err = fmt.Errorf("error: add user_ext")
			return err
		}

		addAssociationErr := repository.AddAssociation(user, userExt)
		if addAssociationErr != nil {
			err = fmt.Errorf("error: update user_association")
			return err
		}

		return err
	})

	return err
}
