package main

import (
	"net/http"
	"xoho-go/controller"
	"xoho-go/database"
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

func updatePassword(c echo.Context) error {
	updatePasswordRes := json.UpdatePasswordRes{
		Status: true,
		Code:   enum.NotEqualOldNew,
	}
	return c.JSON(http.StatusOK, updatePasswordRes)
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
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.GET("/theaters", getTheater)
	e.POST("/signup", signUp)
	e.POST("/login", login)
	e.POST("/update_password", updatePassword)
	e.Logger.Fatal(e.Start(":3000"))
}
