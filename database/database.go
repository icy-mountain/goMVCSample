package database

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var sqlDB *sql.DB

func Init(reset bool, models ...interface{}) error {
	dsn := "root:password@tcp(127.0.0.1:3306)/go_sample"
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err = db.DB()
	if err != nil {
		return err
	}
	if reset {
        db.Migrator().DropTable(models...)
    }
	err = db.Migrator().CreateTable(models...)
	if err != nil {
		return err
	}
	return err
}

func Close() {
	sqlDB.Close()
}

func GetDB() *gorm.DB {
	return db
}
