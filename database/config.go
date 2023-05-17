package database

import "fmt"

type Config struct {
	Host     string
	User     string
	Password string
	DB       string
}

func GetConnectionString(config Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true",
		config.User,
		config.Password,
		config.Host,
		config.DB,
	)
}
