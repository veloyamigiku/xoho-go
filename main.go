package main

import (
	"net/http"
	"xoho-go/database"

	"github.com/labstack/echo/v4"
)

type AreaData struct {
	Name string `json:"name"`
	Sub  string `json:"sub"`
}

type Area struct {
	Id   int    `json:"id" param:"id"`
	Name string `json:"name"`
	Sub  string `json:"sub"`
}

func getAreaDatas(c echo.Context) error {
	areas := []Area{}
	database.DB.Find(&areas)
	area_datas := []AreaData{}
	for _, area := range areas {
		area_data := AreaData{
			Name: area.Name,
			Sub:  area.Sub}
		area_datas = append(
			area_datas,
			area_data)
	}
	return c.JSON(http.StatusOK, area_datas)
}

func getArea(c echo.Context) error {
	area := Area{}
	if err := c.Bind(&area); err != nil {
		return err
	}
	database.DB.Take(&area)
	return c.JSON(http.StatusOK, area)
}

/*
func hello(c echo.Context) error {
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()
	err := sqlDB.Ping()
	if err != nil {
		return c.String(http.StatusInternalServerError, "データベース接続失敗")
	} else {
		return c.String(http.StatusOK, "Hello, World")
	}
}
*/

func main() {
	e := echo.New()
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.GET("/areas", getAreaDatas)
	e.GET("/area/:id", getArea)
	e.Logger.Fatal(e.Start(":3000"))
}
