package controllers

import (
	"icy-mountain/models"
	"github.com/labstack/echo/v4"
)

type PhoneController struct {}

func NewPhoneController() *PhoneController {
	return new(PhoneController)
}

func (phc *PhoneController) PHCHandler(c echo.Context) error {
	ph := new(models.Phone)
	err := echo.QueryParamsBinder(c).
		String("maker", &ph.Maker).
		String("machine", &ph.Machine).
		String("os_version", &ph.OS_version).
		String("color", &ph.Color).
		Time("released", &ph.Released, "2006-01-02").
		Uint("storage", &ph.Storage).
		Uint("price", &ph.Price).
		BindError()
	if err != nil {
		return err
	}
	if err = ph.Create(); err != nil {
		return err
	}
	return err
}
