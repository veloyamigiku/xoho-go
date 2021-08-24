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
