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
