package model

func init() {
	db.AutoMigrate(&TodoModel{}, &UserModel{})
}
