package service

import (
	"crypto/sha256"
	"fmt"
	"xoho-go/database"
	"xoho-go/err"
	"xoho-go/model/db"
	"xoho-go/model/json"
	"xoho-go/model/json/enum"
	"xoho-go/model/repository"
	"xoho-go/utils"

	"gorm.io/gorm"
)

func UpdatePassword(updatePassword json.UpdatePassword) (e error) {

	e = database.DB.Transaction(func(tx *gorm.DB) (txErr error) {

		user, findUserErr := repository.FindUserWithId(updatePassword.UserId)
		if findUserErr != nil {
			txErr = &err.UpdatePasswordError{
				Message: "error: user not exists",
				Code:    enum.NotExistsUser,
			}
			return txErr
		}

		hashedOldPassword := utils.HashPassword(updatePassword.OldPassword)
		if hashedOldPassword != user.Password {
			txErr = &err.UpdatePasswordError{
				Message: "error: old password not equal",
				Code:    enum.NotEqualOld,
			}
			return txErr
		}

		hashedNewPassword := utils.HashPassword(updatePassword.NewPassword)
		if hashedOldPassword == hashedNewPassword {
			txErr = &err.UpdatePasswordError{
				Message: "error: old/new password equal",
				Code:    enum.EqualOldNew,
			}
			return txErr
		}

		updatePasswordErr := repository.UpdatePassword(
			&user,
			hashedNewPassword)
		if updatePasswordErr != nil {
			txErr = &err.UpdatePasswordError{
				Message: "error: update password",
				Code:    enum.UpdateError,
			}
			return txErr
		}

		updateAuthMissCountErr := repository.UpdateAuthMissCount(
			&(user.UserExt),
			0)
		if updateAuthMissCountErr != nil {
			txErr = &err.UpdatePasswordError{
				Message: "error: update user(auth_miss_count)",
				Code:    enum.UpdateError,
			}
			return txErr
		}

		return txErr
	})

	return e
}

func Resign(resign json.Resign) (e error) {

	e = database.DB.Transaction(func(tx *gorm.DB) (txErr error) {

		user, findUserErr := repository.FindUserWithId(resign.UserId)
		if findUserErr != nil {
			txErr = &err.ResignError{
				Code:    enum.ResignCodeNotExistsUser,
				Message: "error: not exists user",
			}
			return txErr
		}

		deleteUserExtErr := repository.DeleteUserExt(&user.UserExt)
		if deleteUserExtErr != nil {
			txErr = &err.ResignError{
				Code:    enum.ResignCodeDbError,
				Message: "error: delete user_ext",
			}
			return txErr
		}

		deleteUserErr := repository.DeleteUser(&user)
		if deleteUserErr != nil {
			txErr = &err.ResignError{
				Code:    enum.ResignCodeDbError,
				Message: "error: delete user",
			}
			return txErr
		}

		return txErr
	})

	return e
}

func Login(login json.Login) (e error) {

	e = database.DB.Transaction(func(tx *gorm.DB) (txErr error) {

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
	return e
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
