package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "admin:12345678@tcp(127.0.0.1:3306)/atomica_blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
