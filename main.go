package main

import (
	"net/http"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/labstack/echo/v4"
)

type Phone struct {
	Maker		string		`json:"maker" form:"maker" query:"maker"`
	Machine		string		`json:"machine" form:"machine" query:"machine"`
	OS_version	string		`json:"os_version" form:"os_version" query:"os_version"`
	Color		string		`json:"color" form:"color" query:"color"`
	Released	time.Time	`json:"released" form:"released" query:"released"`
	Storage		uint		`json:"storage" form:"storage" query:"storage"`
	Price		uint		`json:"price" form:"price" query:"price"`
}

func db_connect() (*gorm.DB, error) {
	message := "[DB connected!]"
	dsn := "root:password@tcp(127.0.0.1:3306)/go_sample"

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		message = "[In db_connect() gorm.Open() error] " + err.Error()
		log.Println(message)
		return nil, err
	}
	log.Println(message)
	return db, err
}

func table_make() error {
	message := "[table make Success!]"
	db , err := db_connect()
	if err != nil {
		message = "[In insert_row() db_connect() error] " + err.Error()
		log.Println(message)
		return err
	}
	sqldb, err := db.DB()
	if err != nil {
		message = "[In insert_row() db.DB() error] " + err.Error()
		log.Println(message)
		return err
	}
	defer sqldb.Close()

	err = db.Migrator().CreateTable(&Phone{})
	if err != nil {
		message = "[In table_make() db.Migrator().CreateTable() error] " + err.Error()
		log.Println(message)
		return err
	}
	log.Println(message)
	return err
}


func Model(c echo.Context) error {
	message := "[Register Success!]"
	db , err := db_connect()
	if err != nil {
		log.Println("[In insert_row() db_connect() error] " + err.Error())
		return err
	}
	sqldb, err := db.DB()
	if err != nil {
		log.Println("[In insert_row() db.DB() error] " + err.Error())
		return err
	}
	defer sqldb.Close()
	
	ph := new(Phone)
	err = echo.QueryParamsBinder(c).
		String("maker", &ph.Maker).
		String("machine", &ph.Machine).
		String("os_version", &ph.OS_version).
		String("color", &ph.Color).
		Time("released", &ph.Released, "2006-01-02").
		Uint("storage", &ph.Storage).
		Uint("price", &ph.Price).
		BindError()
	if err != nil {
		log.Println("[In insert_row() echo.QueryParamsBinder error] " + err.Error())
		return err
	}
	log.Println(ph.Released.Format("2006-01-02"))
	log.Println(ph)
	// err = table_make()
	// if err != nil {
	// 	log.Println("[In insert_row() table_make() error] " + err.Error())
	// 	return err
	// }
	// result := db.Create(&ph)
	// if result.Error != nil {
	// 	log.Println("[In insert_row() db.Create error] " + result.Error.Error())
	// 	return result.Error
	// }
}

func Controller(c echo.Context) error {
	if err := Model(c); err {
		return err
	}
	return nil
}

func HandlerSample(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello!")
}

func main(){
	e := echo.New()
	port := ":8080"
	e.GET("/", HandlerSample)
	e.POST("/register",Controller)

	e.Logger.Fatal(e.Start(port))
}
