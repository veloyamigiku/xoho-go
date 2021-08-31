package service

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"xoho-go/database"
	"xoho-go/err"
	"xoho-go/model/db"
	"xoho-go/model/json"
	"xoho-go/model/json/enum"
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

	t.Run("UpdatePassword", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var user1 db.User
			firstTx1 := tx.
				Debug().
				First(&user1)
			assert.NotNil(t, firstTx1.Error)

			signup := json.Signup{
				Name:     "name1",
				Password: "password1"}
			newUserId, signupErr := _signup2(tx, signup)
			assert.Nil(t, signupErr)
			fmt.Printf("newUserId=%d", newUserId)

			var user2 db.User
			firstTx2 := tx.
				Debug().
				First(&user2)
			assert.Nil(t, firstTx2.Error)

			updatePassword := json.UpdatePassword{
				UserId:      newUserId,
				OldPassword: "password1",
				NewPassword: "password2"}
			_, updatePasswordErr := _updatePassword2(tx, updatePassword)
			assert.Nil(t, updatePasswordErr)

			var user3 db.User
			firstTx3 := tx.
				Debug().
				First(&user3)
			assert.Nil(t, firstTx3.Error)

			assert.Equal(
				t,
				utils.HashPassword("password2"),
				user3.Password)
			assert.Equal(
				t,
				0,
				user3.UserExt.AuthMissCount)

			return errors.New("rollback UpdatePassword")
		})

	})

	t.Run("UpdatePassword_OldPasswordNotEqual", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var user1 db.User
			firstTx1 := tx.
				Debug().
				First(&user1)
			assert.NotNil(t, firstTx1.Error)

			signup := json.Signup{
				Name:     "name1",
				Password: "password1"}
			newUserId, signupErr := _signup2(tx, signup)
			assert.Nil(t, signupErr)
			fmt.Printf("newUserId=%d", newUserId)

			var user2 db.User
			firstTx2 := tx.
				Debug().
				First(&user2)
			assert.Nil(t, firstTx2.Error)

			updatePassword := json.UpdatePassword{
				UserId:      newUserId,
				OldPassword: "password2",
				NewPassword: "password3"}
			_, updatePasswordErr := _updatePassword2(tx, updatePassword)
			assert.NotNil(t, updatePasswordErr)

			if updatePasswordError, ok := updatePasswordErr.(*err.UpdatePasswordError); ok {
				assert.Equal(
					t,
					enum.NotEqualOld,
					updatePasswordError.Code)
			}

			return errors.New("rollback UpdatePassword_OldPasswordNotEqual")

		})

	})

	t.Run("UpdatePassword_OldNewPasswordEqual", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var user1 db.User
			firstTx1 := tx.
				Debug().
				First(&user1)
			assert.NotNil(t, firstTx1.Error)

			signup := json.Signup{
				Name:     "name1",
				Password: "password1"}
			newUserId, signupErr := _signup2(tx, signup)
			assert.Nil(t, signupErr)
			fmt.Printf("newUserId=%d", newUserId)

			var user2 db.User
			firstTx2 := tx.
				Debug().
				First(&user2)
			assert.Nil(t, firstTx2.Error)

			updatePassword := json.UpdatePassword{
				UserId:      newUserId,
				OldPassword: "password1",
				NewPassword: "password1"}
			_, updatePasswordErr := _updatePassword2(tx, updatePassword)
			assert.NotNil(t, updatePasswordErr)

			if updatePasswordError, ok := updatePasswordErr.(*err.UpdatePasswordError); ok {
				assert.Equal(
					t,
					enum.EqualOldNew,
					updatePasswordError.Code)
			}

			return errors.New("rollback UpdatePassword_OldNewPasswordEqual")

		})

	})

	t.Run("UpdatePassword_UserNotExists", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			var user1 db.User
			firstTx1 := tx.
				Debug().
				First(&user1)
			assert.NotNil(t, firstTx1.Error)

			updatePassword := json.UpdatePassword{
				UserId:      0,
				OldPassword: "password1",
				NewPassword: "password2"}
			_, updatePasswordErr := _updatePassword2(tx, updatePassword)
			assert.NotNil(t, updatePasswordErr)

			if updatePasswordError, ok := updatePasswordErr.(*err.UpdatePasswordError); ok {
				assert.Equal(
					t,
					enum.NotExistsUser,
					updatePasswordError.Code)
			}

			return errors.New("rollback UpdatePassword_UserNotExists")
		})

	})

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
