package main

import (
	"net/http"
	"xoho-go/database"

	"github.com/labstack/echo/v4"

	. "xoho-go/model"
)

func getTheaters(c echo.Context) error {
	theaters := []Theater{}
	database.DB.Debug().Model(&theaters).Preload("Type").Preload("Area").Preload("Prefecture").Find(&theaters)
	return c.JSON(http.StatusOK, theaters)
}

func main() {
	e := echo.New()
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.GET("/theaters", getTheaters)
	e.Logger.Fatal(e.Start(":3000"))
}
