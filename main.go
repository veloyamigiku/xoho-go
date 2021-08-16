package main

import (
	"net/http"
	"os"
	"xoho-go/controller"
	"xoho-go/database"
	"xoho-go/err"
	"xoho-go/model/json"
	"xoho-go/model/json/enum"
	"xoho-go/service"

	"github.com/labstack/echo/v4"
)

func getTheater(c echo.Context) error {

	// 劇場情報取得時のクエリ文字列をパースする。
	var queryTheater controller.QueryTheater
	if err := c.Bind(&queryTheater); err != nil {
		panic(err)
	}
	res := []json.TheaterRes{}

	// クエリ文字列の項目(type)で条件分岐する。
	switch queryTheater.Type {
	case "all":
		res = service.GetAllTheaters()
	case "all_type":
		res = service.GetAllTypeTheaters()
	}
	return c.JSON(http.StatusOK, res)
}

func updatePassword(c echo.Context) (e error) {

	updatePasswordRes := json.UpdatePasswordRes{
		Status: true,
		Code:   enum.NoError,
		Msg:    "",
	}
	var updatePassword json.UpdatePassword
	if e = c.Bind(&updatePassword); e != nil {
		updatePasswordRes.Status = false
		updatePasswordRes.Code = enum.ParseParamError
		updatePasswordRes.Msg = "error: parse param"
		return c.JSON(http.StatusOK, updatePasswordRes)
	}

	e = service.UpdatePassword(updatePassword)
	if e != nil {
		updatePasswordRes.Status = false
		if updatePasswordError, ok := e.(*err.UpdatePasswordError); ok {
			updatePasswordRes.Code = updatePasswordError.Code
			updatePasswordRes.Msg = updatePasswordError.Message
		}
		return c.JSON(http.StatusOK, updatePasswordRes)
	}

	return c.JSON(http.StatusOK, updatePasswordRes)
}

func resign(c echo.Context) error {

	resignRes := json.ResignRes{
		Status: true,
		Code:   enum.ResignCodeNoError,
		Msg:    "",
	}

	var resign json.Resign
	if parseErr := c.Bind(&resign); parseErr != nil {
		resignRes.Status = false
		resignRes.Code = enum.ResignCodeParseParamError
		resignRes.Msg = "error: parse param."
		return c.JSON(http.StatusOK, resignRes)
	}

	serivceErr := service.Resign(resign)
	if serivceErr != nil {
		resignRes.Status = false
		if resignErr, ok := serivceErr.(*err.ResignError); ok {
			resignRes.Code = resignErr.Code
			resignRes.Msg = resignErr.Message
			return c.JSON(http.StatusOK, resignRes)
		}
	}

	return c.JSON(http.StatusOK, resignRes)
}

func login(c echo.Context) error {

	loginRes := json.LoginRes{
		Status: true,
	}
	var login json.Login
	if err := c.Bind(&login); err != nil {
		loginRes.Status = false
		return c.JSON(http.StatusOK, loginRes)
	}

	err := service.Login(login)
	if err != nil {
		loginRes.Status = false
		return c.JSON(http.StatusOK, loginRes)
	}

	return c.JSON(http.StatusOK, loginRes)
}

func signUp(c echo.Context) error {

	// POSTデータをパースする。
	signupRes := json.SignupRes{
		Status: true,
	}
	var signup json.Signup
	if err := c.Bind(&signup); err != nil {
		signupRes.Status = false
		return c.JSON(http.StatusOK, signupRes)
	}

	err := service.SignUp(signup)
	if err != nil {
		signupRes.Status = false
		return c.JSON(http.StatusOK, signupRes)
	}

	return c.JSON(http.StatusOK, signupRes)
}

func main() {
	e := echo.New()
	curDir, _ := os.Getwd()
	database.Connect(curDir)
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.GET("/theaters", getTheater)
	e.POST("/signup", signUp)
	e.POST("/login", login)
	e.POST("/update_password", updatePassword)
	e.POST("/resign", resign)
	e.Logger.Fatal(e.Start(":3000"))
}
