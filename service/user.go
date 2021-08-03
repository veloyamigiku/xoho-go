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

func updatePassword(updatePassword json.UpdatePassword) (err error) {
	return nil
}

func Login(login json.Login) (err error) {

	err = database.DB.Transaction(func(tx *gorm.DB) (txErr error) {

		user, findUserErr := repository.FindUserWithName(login.Name)
		if findUserErr != nil {
			txErr = fmt.Errorf("error: find user")
			return txErr
		}

		if user.UserExt.AuthMissCount >= 3 {
			txErr = fmt.Errorf("error: user frozen")
			return txErr
		}

		hash := utils.HashPassword(login.Password)
		trueHash := user.Password
		if hash == trueHash {
			updateAuthMissCountErr := repository.UpdateAuthMissCount(
				&(user.UserExt),
				0)
			if updateAuthMissCountErr != nil {
				txErr = fmt.Errorf("error: update user(auth_miss_count)")
				return txErr
			}
		} else {
			updateAuthMissCountErr := repository.UpdateAuthMissCount(
				&(user.UserExt),
				user.UserExt.AuthMissCount+1)
			if updateAuthMissCountErr != nil {
				txErr = fmt.Errorf("error: update user(auth_miss_count)")
				return txErr
			}
			txErr = fmt.Errorf("error: not equal password")
			return txErr
		}
		return txErr
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
