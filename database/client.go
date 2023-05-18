package database

import (
	user_entity "akbariskndr/todo-service-gin/modules/auth/entity"
	todo_entity "akbariskndr/todo-service-gin/modules/todo/entity"
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
	Connector.AutoMigrate(&todo_entity.TodoEntity{})
	Connector.AutoMigrate(&user_entity.UserEntity{})
	log.Println("Connected to database")

	return nil
}
