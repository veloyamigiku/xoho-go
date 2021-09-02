package service

import (
	"crypto/sha256"
	"errors"
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

func _updatePassword2(
	tx *gorm.DB,
	updatePassword json.UpdatePassword) (int, error) {

	var e error
	e = nil

	user, findUserErr := repository.FindUserWithId(
		tx,
		updatePassword.UserId)
	if findUserErr != nil {
		e = &err.UpdatePasswordError{
			Message: "error: user not exists",
			Code:    enum.NotExistsUser,
		}
		return -1, e
	}

	hashedOldPassword := utils.HashPassword(updatePassword.OldPassword)
	if hashedOldPassword != user.Password {
		e = &err.UpdatePasswordError{
			Message: "error: old password not equal",
			Code:    enum.NotEqualOld,
		}
		return -1, e
	}

	hashedNewPassword := utils.HashPassword(updatePassword.NewPassword)
	if hashedOldPassword == hashedNewPassword {
		e = &err.UpdatePasswordError{
			Message: "error: old/new password equal",
			Code:    enum.EqualOldNew,
		}
		return -1, e
	}

	updatePasswordErr := repository.UpdatePassword(
		tx,
		&user,
		hashedNewPassword)
	if updatePasswordErr != nil {
		e = &err.UpdatePasswordError{
			Message: "error: update password",
			Code:    enum.UpdateError,
		}
		return -1, e
	}

	updateAuthMissCountErr := repository.UpdateAuthMissCount(
		tx,
		&(user.UserExt),
		0)
	if updateAuthMissCountErr != nil {
		e = &err.UpdatePasswordError{
			Message: "error: update user(auth_miss_count)",
			Code:    enum.UpdateError,
		}
		return -1, e
	}

	return user.Id, e
}

func UpdatePassword2(updatePassword json.UpdatePassword) (int, error) {
	var user_id int
	var orgErr error
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		tmp_user_id, e := _updatePassword2(tx, updatePassword)
		user_id = tmp_user_id
		orgErr = e
		if utils.IsTest() {
			return errors.New("rollback UpdatePassword for Test")
		}
		return e
	})
	if utils.IsTest() {
		return user_id, orgErr
	}
	return user_id, err
}

func UpdatePassword(updatePassword json.UpdatePassword) (e error) {

	e = database.DB.Transaction(func(tx *gorm.DB) (txErr error) {

		user, findUserErr := repository.FindUserWithId(
			tx,
			updatePassword.UserId)
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
			tx,
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
			tx,
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

		user, findUserErr := repository.FindUserWithId(
			tx,
			resign.UserId)
		if findUserErr != nil {
			txErr = &err.ResignError{
				Code:    enum.ResignCodeNotExistsUser,
				Message: "error: not exists user",
			}
			return txErr
		}

		deleteUserExtErr := repository.DeleteUserExt(
			tx,
			&user.UserExt)
		if deleteUserExtErr != nil {
			txErr = &err.ResignError{
				Code:    enum.ResignCodeDbError,
				Message: "error: delete user_ext",
			}
			return txErr
		}

		_, deleteUserErr := repository.DeleteUser(
			tx,
			&user)
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

		user, findUserErr := repository.FindUserWithName(
			tx,
			login.Name)
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
				tx,
				&(user.UserExt),
				0)
			if updateAuthMissCountErr != nil {
				txErr = fmt.Errorf("error: update user(auth_miss_count)")
				return txErr
			}
		} else {
			updateAuthMissCountErr := repository.UpdateAuthMissCount(
				tx,
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

func SignUp2(signup json.Signup) (int, error) {
	var user_id int
	var orgErr error
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		tmp_user_id, e := _signup2(tx, signup)
		user_id = tmp_user_id
		orgErr = e
		if utils.IsTest() {
			return errors.New("rollback Signup2 for Test")
		}
		return e
	})
	if utils.IsTest() {
		return user_id, orgErr
	}
	return user_id, err
}

func _signup2(tx *gorm.DB, signup json.Signup) (int, error) {

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

	exists, existsUserErr := repository.ExistsUser(
		tx,
		user)
	if existsUserErr != nil || exists {
		err = fmt.Errorf("error: exists user")
		return -1, err
	}

	addUserErr := repository.AddUser(tx, user)
	if addUserErr != nil {
		err = fmt.Errorf("error: add user")
		return -1, err
	}

	addUserExtErr := repository.AddUserExt(tx, userExt)
	if addUserExtErr != nil {
		err = fmt.Errorf("error: add user_ext")
		return -1, err
	}

	addAssociationErr := repository.AddAssociation(tx, user, userExt)
	if addAssociationErr != nil {
		err = fmt.Errorf("error: update user_association")
		return -1, err
	}

	return user.Id, err

}
