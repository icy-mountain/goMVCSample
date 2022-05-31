// Package database
// USE: importしてdatabase.Init()した後に、foo := database.Get()すれば使える。
package database

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// sqlDBはデータベース接続を閉じるためだけに使用される。
var db *gorm.DB
var sqlDB *sql.DB

// Init resetがTrueなら、既にデータベースに存在するテーブルを削除して新しいテーブルを作成する。
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
        err = db.Migrator().DropTable(models...)
    }
	err = db.Migrator().CreateTable(models...)
	if err != nil {
		return err
	}
	return err
}

func Close() {
	err := sqlDB.Close()
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
