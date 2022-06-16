package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	dsn := "host=tai.db.elephantsql.com user=hnwhqtul password=sub7Txm66RYypQFNO5ChapHVgXPcNCXV dbname=hnwhqtul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
