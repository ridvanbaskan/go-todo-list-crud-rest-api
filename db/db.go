package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func InitDB(dst ...interface{}) {
	dsn := "host=localhost port=5432 dbname=postgres user=postgres password=123654 sslmode=prefer connect_timeout=10"
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("db successfully init")
	}

	Db.AutoMigrate(dst...)

}
