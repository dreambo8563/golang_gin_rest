package main

import (
	"vincent.com/golangginrest/router"

	"vincent.com/golangginrest/db"
)

func init() {
	db.InitDB()
}

func main() {
	router.InitRouter()
}
