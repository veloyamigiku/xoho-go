package service

import (
	"fmt"
	"xoho-go/model/db"
	"xoho-go/model/json"
	"xoho-go/model/repository"
)

func SignUp(signup json.Signup) error {

	// Todo
	// ユーザ名の唯一性ををサポートする。
	// トランザクション処理に書き換える。
	userExt := &db.UserExt{
		AuthMissCount: 0,
	}
	user := &db.User{
		Name:     signup.Name,
		Password: signup.Password,
		UserExt:  *userExt,
	}
	var err error
	err = nil
	addUserErr := repository.AddUser(user)
	if addUserErr != nil {
		err = fmt.Errorf("hogehoge1")
		return err
	}
	addUserExtErr := repository.AddUserExt(userExt)
	if addUserExtErr != nil {
		err = fmt.Errorf("hogehoge2")
		return err
	}
	addAssociationErr := repository.AddAssociation(user, userExt)
	if addAssociationErr != nil {
		err = fmt.Errorf("hogehoge3")
		return err
	}
	return err
}
