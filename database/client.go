package database

import (
	"akbariskndr/todo-service-gin/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open(mysql.New(mysql.Config{
		DSN: connectionString,
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	Connector.AutoMigrate(&entity.Todo{})
	log.Println("Connected to database")

	return nil
}
