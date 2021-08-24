package service

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"xoho-go/database"
	"xoho-go/model/db"
	"xoho-go/model/json"
	"xoho-go/utils"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUser(t *testing.T) {

	curDir, _ := os.Getwd()
	envDir := fmt.Sprintf("%s/..", curDir)
	database.Connect(envDir)
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	t.Run("SignUp", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			signup := json.Signup{
				Name:     "velo.yamigiku@gmail.com",
				Password: "velo.yamigiku@gmail.com",
			}
			user_id, signupErr := _signup2(tx, signup)
			assert.Nil(t, signupErr)

			var newUser db.User
			firstErr1 := tx.
				Debug().
				First(&newUser, user_id)
			assert.Nil(t, firstErr1.Error)
			fmt.Print(newUser)

			assert.Equal(
				t,
				user_id,
				newUser.Id)
			assert.Equal(
				t,
				signup.Name,
				newUser.Name)
			assert.Equal(
				t,
				utils.HashPassword(signup.Password),
				newUser.Password)
			assert.Equal(
				t,
				0,
				newUser.UserExt.AuthMissCount)

			return errors.New("rollback Signup")
		})
	})

}
