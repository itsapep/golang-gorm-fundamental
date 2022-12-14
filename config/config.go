package config

import (
	"fmt"
	"golang-gorm-fundamental/utils"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	dbHost string
	dbPort string
	dbUser string
	dbPass string
	dbName string
}

type Config struct {
	db *gorm.DB
}

func (c *Config) initDb() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dbConfig := DBConfig{
		dbHost, dbPort, dbUser, dbPass, dbName,
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbConfig.dbHost, dbConfig.dbUser, dbConfig.dbPass, dbConfig.dbName, dbConfig.dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.IsError(err)

	c.db = db
}

func (c *Config) DBConn() *gorm.DB {
	fmt.Println("connected")
	return c.db
}

func (c *Config) DBClose() error {
	db, err := c.db.DB()
	utils.IsError(err)
	err = db.Close()
	utils.IsError(err)
	return err
}

func NewConfig() Config {
	cfg := Config{}
	cfg.initDb()
	return cfg
}
