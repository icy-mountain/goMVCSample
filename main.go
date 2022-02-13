package main

import (
	"net/http"

	"icy-mountain/models"
	"icy-mountain/controllers"
	"icy-mountain/database"
	"github.com/labstack/echo/v4"
)

func HandlerSample(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello!")
}

func main(){
	e := echo.New()
	port := ":8080"
	phc := controllers.NewPhoneController()
	err := database.Init(true, new(models.Phone))
	if err != nil {
		panic(err)
	}
    defer database.Close()

	e.GET("/", HandlerSample)
	e.POST("/register", phc.PHCHandler)

	e.Logger.Fatal(e.Start(port))
}
