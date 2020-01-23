package model

import (
	"github.com/jinzhu/gorm"

	"log"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init() {
	db := getConnect()
	defer db.Close()
	db.LogMode(true)
	if !db.HasTable(&Source{}) {
		db.CreateTable(&Source{})
	}

	if !db.HasTable(&Subscribe{}) {
		db.CreateTable(&Subscribe{})
	}

	if !db.HasTable(&Content{}) {
		db.CreateTable(&Content{})
	}
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
}

func getConnect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./data.db")
	if err != nil {
		log.Println(err.Error())
		panic("连接SQLite数据库失败")
	}
	return db
}

//EditTime timestamp
type EditTime struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
