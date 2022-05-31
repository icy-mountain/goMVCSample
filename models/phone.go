// Package models
/*
TODO:疎結合なデータベース構造への変更。
*/
package models

import (
	"icy-mountain/database"
	"time"

	"gorm.io/gorm"
)

// Phone 仕様書ではIDについて言及がなかったが追加。
type Phone struct {
	gorm.Model
	ID			uint		`json:"id" gorm:"primaryKey not null autoIncrement"`
	Maker		string		`json:"maker" form:"maker" query:"maker"`
	Machine		string		`json:"machine" form:"machine" query:"machine"`
	OSVersion	string		`json:"os_version" form:"os_version" query:"os_version"`
	Color		string		`json:"color" form:"color" query:"color"`
	Released	time.Time	`json:"released" form:"released" query:"released"`
	Storage		uint		`json:"storage" form:"storage" query:"storage"`
	Price		uint		`json:"price" form:"price" query:"price"`
}

func (ph *Phone) Create() (err error) {
	db := database.GetDB()
	return db.Create(ph).Error
}
