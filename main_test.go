package main

import (
	js "encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"xoho-go/database"
	"xoho-go/model/json"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	curDir, _ := os.Getwd()
	database.Connect(curDir)
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	t.Run("SignUp", func(t *testing.T) {
		e := echo.New()
		signup := json.Signup{
			Name:     "name1",
			Password: "pass1",
		}
		s, _ := js.Marshal(signup)
		req := httptest.NewRequest(
			http.MethodPost,
			"/",
			strings.NewReader(string(s)))
		req.Header.Set(
			echo.HeaderContentType,
			echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, signUp(c)) {
			assert.Equal(
				t,
				http.StatusOK,
				rec.Code)
		}

	})

}
