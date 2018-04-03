package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"vincent.com/golangginrest/service/persistdata"
	"vincent.com/golangginrest/utils/logger"
)

var (
	db  = mysql.DB
	log = logger.Log()
)

// BasicModel is the basic struct for each model
type BasicModel struct {
	ID        uint       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

// DB return the single instance
func DB() *gorm.DB {
	return db
}
