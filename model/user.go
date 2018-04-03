package model

import (
	"github.com/jinzhu/gorm"
)

type (
	// UserModel - user info
	UserModel struct {
		BasicModel
		Phone    string `gorm:"type:varchar(13);unique_index;default:'';NOT NULL" json:"phone" binding:"required,len=13"`
		Password string `gorm:"type:varchar(16);default:'';NOT NULL" json:"password" binding:"required,max=16,min=6"`
	}
)

// TableName custormized table name
func (UserModel) TableName() string {
	return "profiles"
}

// New will create user recor
func (user *UserModel) New() error {
	return db.Create(&user).Error
}

// FindByInfo - will check whether the user exist by the phone and password info
func FindByInfo(phone, password string) (UserModel, bool) {
	userInfo := UserModel{}
	err := db.Select("id").Where(&UserModel{Phone: phone, Password: password}).Find(&userInfo).Error
	if gorm.IsRecordNotFoundError(err) {
		return userInfo, false
	}
	if err != nil {
		log.Errorln(err.Error())
		return userInfo, false
	}
	return userInfo, true
}
