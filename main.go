package main

import (
	"net/http"
	"xoho-go/controller"
	"xoho-go/database"
	"xoho-go/model/json"
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

func main() {
	e := echo.New()
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.GET("/theaters", getTheater)
	e.Logger.Fatal(e.Start(":3000"))
}
