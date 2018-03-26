package dao

import (
	"strconv"
	"strings"

	"vincent.com/golangginrest/utils/logger"

	"github.com/jinzhu/gorm"
	"vincent.com/golangginrest/config"
	"vincent.com/golangginrest/model"

	// import to init mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB instantce singleton
var (
	DB  *gorm.DB
	err error
	log = logger.Log()
)

// Init - connect db and AutoMigrate
func init() {
	//open a db connection
	dialects := dialectsString()

	DB, err = gorm.Open("mysql", dialects)

	if err != nil {
		log.Panicln("connected to db failed", err)
	}
	DB.AutoMigrate(&model.TodoModel{})
}

func dialectsString() string {
	var config = config.Item()
	var b strings.Builder
	b.WriteString(config.Mysql.Username)
	b.WriteString(":")
	b.WriteString(config.Mysql.Password)
	b.WriteString("@tcp(")
	b.WriteString(config.Mysql.Host)
	b.WriteString(":")
	b.WriteString(strconv.Itoa(config.Mysql.Port))
	b.WriteString(")/")
	b.WriteString(config.Mysql.Dbname)
	return b.String()
}
