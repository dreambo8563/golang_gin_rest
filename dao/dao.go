package dao

import (
	"github.com/jinzhu/gorm"
	"vincent.com/golangginrest/service/persistdata"
)

var db = mysql.DB

// DB return the single instance
func DB() *gorm.DB {
	return db
}
