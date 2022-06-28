package main

import (
	"golang-gorm-fundamental/config"
	"golang-gorm-fundamental/model/entity"
	"golang-gorm-fundamental/utils"
)

func main() {
	// dbHost := "localhost"
	// dbPort := "5432"
	// dbUser := "postgres"
	// dbPass := "12345678"
	// dbName := "db_enigma_shop_v2"

	cfg := config.NewConfig()
	db := cfg.DBConn()
	defer cfg.DBClose()

	err := db.AutoMigrate(&entity.Customer{})
	utils.IsError(err)
}
