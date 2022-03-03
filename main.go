/*
ROLE:URLルーティングとサーバー建てを行う。
TODO:ルーティングを別ファイルで行う。複数の役割を一つのファイルで行うのは好ましくないため。
*/
package main

import (
	"net/http"

	"icy-mountain/models"
	"icy-mountain/controllers"
	"icy-mountain/database"
	"github.com/labstack/echo/v4"
)

// 「/」へGETが要求された時に動く。単なる動作確認用。別に消して良い
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
