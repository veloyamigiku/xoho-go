package repository

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"xoho-go/database"
	"xoho-go/model/db"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUser(t *testing.T) {

	curDir, _ := os.Getwd()
	envDir := fmt.Sprintf("%s/../..", curDir)
	database.Connect(envDir)
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	t.Run("ExistsUser_exists", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var user1 db.User
			firstUserTx1 := tx.
				Debug().
				First(&user1)
			assert.NotNil(t, firstUserTx1.Error)

			newUser := db.User{
				Name:     "name1",
				Password: "password1"}
			addUserErr := AddUser(tx, &newUser)
			assert.Nil(t, addUserErr)

			var user2 db.User
			firstUserTx2 := tx.
				Debug().
				First(&user2)
			assert.Nil(t, firstUserTx2.Error)

			exists, existsUserErr := ExistsUser(
				tx,
				&user2)
			assert.Nil(
				t,
				existsUserErr)
			assert.Equal(
				t,
				true,
				exists)

			return errors.New("rollback ExistsUser_exists")
		})

	})

	t.Run("ExistsUser_not_exists", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var user1 db.User
			firstUserTx1 := tx.
				Debug().
				First(&user1)
			assert.NotNil(t, firstUserTx1.Error)

			exists, existsUserErr := ExistsUser(
				tx,
				&db.User{})
			assert.Nil(
				t,
				existsUserErr)
			assert.Equal(
				t,
				false,
				exists)

			return errors.New("rollback ExistsUser_not_exists")
		})
	})
	t.Run("DeleteUser_exists", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var user1 db.User
			firstUserTx1 := tx.
				Debug().
				First(&user1)
			assert.NotNil(t, firstUserTx1.Error)

			newUser := db.User{
				Name:     "name1",
				Password: "password1"}
			addUserErr := AddUser(tx, &newUser)
			assert.Nil(t, addUserErr)

			var user2 db.User
			firstUserTx2 := tx.
				Debug().
				First(&user2)
			assert.Nil(t, firstUserTx2.Error)

			deleteCount, deleteUserErr := DeleteUser(
				tx,
				&user2)
			assert.Nil(
				t,
				deleteUserErr)
			assert.Equal(
				t,
				int64(1),
				deleteCount)

			var user3 db.User
			firstUserTx3 := tx.
				Debug().
				First(&user3)
			assert.NotNil(t, firstUserTx3.Error)

			return errors.New("rollback DeleteUser_exists")
		})

	})
	t.Run("DeleteUser_not_exists", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var user1 db.User
			firstUserTx1 := tx.
				Debug().
				First(&user1)
			assert.NotNil(t, firstUserTx1.Error)

			deleteCount, deleteUserErr := DeleteUser(
				tx,
				&db.User{})
			assert.Nil(
				t,
				deleteUserErr)
			assert.Equal(
				t,
				int64(0),
				deleteCount)
			return errors.New("rollback DeleteUser_not_exists")

		})
	})

	t.Run("DeleteUserExt_exists", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var userExt1 db.UserExt
			firstUserExtTx1 := tx.
				Debug().
				First(&userExt1)
			assert.NotNil(t, firstUserExtTx1.Error)

			newUserExt := db.UserExt{
				AuthMissCount: 0}
			addUserExtErr := AddUserExt(tx, &newUserExt)
			assert.Nil(t, addUserExtErr)

			var userExt2 db.UserExt
			firstUserExtTx2 := tx.
				Debug().
				First(&userExt2)
			assert.Nil(t, firstUserExtTx2.Error)

			deleteUserExtErr := DeleteUserExt(
				tx,
				&userExt2)
			assert.Nil(
				t,
				deleteUserExtErr)

			var userExt3 db.UserExt
			firstUserExtTx3 := tx.
				Debug().
				First(&userExt3)
			assert.NotNil(t, firstUserExtTx3.Error)

			return errors.New("rollback DeleteUserExt_exists")
		})

	})

	t.Run("DeleteUserExt_not_exists", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var userExt1 db.UserExt
			firstUserExtTx1 := tx.
				Debug().
				First(&userExt1)
			assert.NotNil(t, firstUserExtTx1.Error)

			deleteUserExtErr := DeleteUserExt(
				tx,
				&db.UserExt{})
			assert.Nil(
				t,
				deleteUserExtErr)

			return errors.New("rollback DeleteUserExt_not_exists")
		})

	})

	t.Run("FindUserWithName_exists", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var user1 db.User
			firstUserTx1 := tx.
				Debug().
				First(&user1)
			assert.NotNil(t, firstUserTx1.Error)

			newUser := db.User{
				Name:     "name1",
				Password: "password1"}
			addUserErr := AddUser(tx, &newUser)
			assert.Nil(t, addUserErr)

			findUser, findUserWithNameErr := FindUserWithName(
				tx,
				"name1")
			assert.Nil(
				t,
				findUserWithNameErr)
			assert.Equal(
				t,
				newUser.Name,
				findUser.Name)

			return errors.New("rollback FindUserWithName_exists")
		})

	})
	t.Run("FindUserWithName_not_exists", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var user1 db.User
			firstUserTx1 := tx.
				Debug().
				First(&user1)
			assert.NotNil(t, firstUserTx1.Error)

			_, findUserWithNameErr := FindUserWithName(
				tx,
				"name1")
			assert.NotNil(
				t,
				findUserWithNameErr)

			return errors.New("rollback FindUserWithName_not_exists")
		})

	})
	t.Run("FindUserWithId_exists", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var user1 db.User
			firstUserTx1 := tx.
				Debug().
				First(&user1)
			assert.NotNil(t, firstUserTx1.Error)

			newUser := db.User{
				Name:     "name1",
				Password: "password1"}
			addUserErr := AddUser(tx, &newUser)
			assert.Nil(t, addUserErr)

			findUser, findUserWithIdErr := FindUserWithId(
				tx,
				newUser.Id)
			assert.Nil(
				t,
				findUserWithIdErr)
			assert.Equal(
				t,
				newUser.Id,
				findUser.Id)

			return errors.New("rollback FindUserWithId_exists")
		})

	})

	t.Run("FindUserWithId_not_exists", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var user1 db.User
			firstUserTx1 := tx.
				Debug().
				First(&user1)
			assert.NotNil(t, firstUserTx1.Error)

			_, findUserWithIdErr := FindUserWithId(
				tx,
				1234)
			assert.NotNil(
				t,
				findUserWithIdErr)

			return errors.New("rollback FindUserWithId_not_exists")

		})

	})

	t.Run("AddAssociation", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var user1 db.User
			firstUserTx1 := tx.
				Debug().
				First(&user1)
			assert.NotNil(t, firstUserTx1.Error)

			newUser := db.User{
				Name:     "name1",
				Password: "password1"}
			addUserErr := AddUser(tx, &newUser)
			assert.Nil(t, addUserErr)

			var user2 db.User
			firstUserTx2 := tx.
				Debug().
				First(&user2)
			assert.Nil(t, firstUserTx2.Error)

			var userExt1 db.UserExt
			firstUserExtTx1 := tx.
				Debug().
				First(&userExt1)
			assert.NotNil(t, firstUserExtTx1.Error)

			newUserExt := db.UserExt{
				AuthMissCount: 0}
			addUserExtErr := AddUserExt(tx, &newUserExt)
			assert.Nil(t, addUserExtErr)

			var userExt2 db.UserExt
			firstUserExtTx2 := tx.
				Debug().
				First(&userExt2)
			assert.Nil(t, firstUserExtTx2.Error)

			addAssociationErr := AddAssociation(
				tx,
				&user2,
				&userExt2)
			assert.Nil(t, addAssociationErr)

			user3, findUserWithIdErr := FindUserWithId(
				tx,
				user2.Id)
			assert.Nil(
				t,
				findUserWithIdErr)

			assert.Equal(
				t,
				userExt2.Id,
				user3.UserExt.Id)

			return errors.New("rollback AddAssociate")
		})

	})

	t.Run("UpdateAuthMissCount", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var userExt1 db.UserExt
			firstTx1 := tx.
				Debug().
				First(&userExt1)
			assert.NotNil(t, firstTx1.Error)

			newUserExt := db.UserExt{
				AuthMissCount: 0}
			addUserExtErr := AddUserExt(tx, &newUserExt)
			assert.Nil(t, addUserExtErr)

			var userExt2 db.UserExt
			firstTx2 := tx.
				Debug().
				First(&userExt2)
			assert.Nil(t, firstTx2.Error)

			updateAuthMissCountErr := UpdateAuthMissCount(
				tx,
				&userExt2,
				100)
			assert.Nil(t, updateAuthMissCountErr)

			var userExt3 db.UserExt
			firstTx3 := tx.
				Debug().
				First(&userExt3)
			assert.Nil(t, firstTx3.Error)

			assert.Equal(
				t,
				100,
				userExt3.AuthMissCount)

			return errors.New("rollback UpdateAuthMissCount")
		})

	})

	t.Run("UpdatePassword", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var user1 db.User
			firstTx1 := tx.
				Debug().
				First(&user1)
			assert.NotNil(t, firstTx1.Error)

			newUser := db.User{
				Name:     "name1",
				Password: "password1"}
			addUserErr := AddUser(tx, &newUser)
			assert.Nil(t, addUserErr)

			var user2 db.User
			firstTx2 := tx.
				Debug().
				First(&user2)
			assert.Nil(t, firstTx2.Error)

			updatePasswordErr := UpdatePassword(
				tx,
				&user2,
				"password2")
			assert.Nil(t, updatePasswordErr)

			var user3 db.User
			firstTx3 := tx.
				Debug().
				First(&user3)
			assert.Nil(t, firstTx3.Error)
			assert.Equal(
				t,
				"password2",
				user3.Password)

			return errors.New("rollback UpdatePassword")
		})

	})
	t.Run("AddUser", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {
			var user db.User
			firstTx1 := tx.
				Debug().
				First(&user)
			assert.NotNil(t, firstTx1.Error)

			newUser := db.User{
				Name:     "name1",
				Password: "password1",
			}
			addUserError := AddUser(tx, &newUser)
			assert.Nil(t, addUserError)

			firstTx2 := tx.
				Debug().
				First(&user)
			assert.Nil(t, firstTx2.Error)

			return errors.New("rollback AddUser")
		})
	})

	t.Run("AddUserExt", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var userExt db.UserExt
			firstTx1 := tx.
				Debug().
				First(&userExt)
			assert.NotNil(t, firstTx1.Error)

			newUserExt := db.UserExt{
				Id:            1,
				AuthMissCount: 2,
			}
			addUserExtErr := AddUserExt(tx, &newUserExt)
			assert.Nil(t, addUserExtErr)

			firstTx2 := tx.
				Debug().
				First(&userExt)
			assert.Nil(t, firstTx2.Error)

			return errors.New("rollback AddUserExt")
		})

	})

}
