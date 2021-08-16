package repository

import (
	"fmt"
	"os"
	"testing"
	"xoho-go/database"
	"xoho-go/model/db"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {

	curDir, _ := os.Getwd()
	envDir := fmt.Sprintf("%s/../..", curDir)
	database.Connect(envDir)
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	t.Run("AddUser", func(t *testing.T) {
		var user db.User
		tx := database.
			DB.
			Debug().
			First(&user)
		assert.NotNil(t, tx.Error)
	})

	t.Run("AddUserExt", func(t *testing.T) {
		var userExt db.UserExt
		tx := database.
			DB.
			Debug().
			First(&userExt)
		assert.NotNil(t, tx.Error)
	})
}
