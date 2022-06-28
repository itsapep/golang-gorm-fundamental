package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPass := "12345678"
	dbName := "db_enigma_shop_v2"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	enigmaDB, err := db.DB()
	defer func(enigmaDB *sql.DB) {
		err = enigmaDB.Close()
		if err != nil {
			panic(err)
		}
	}(enigmaDB)

	err = db.AutoMigrate(&Customer{})
	if err != nil {
		panic(err)
	}
}
