package database

import (
	"github.com/jinzhu/gorm"
	"vincent.com/golangginrest/model"

	// import to init mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB instantce singleton
var DB *gorm.DB

// InitDB - connect db and AutoMigrate
func InitDB() {
	//open a db connection
	var err error

	DB, err = gorm.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&model.TodoModel{})
}
