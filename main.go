package main

import (
	"vincent.com/golangginrest/router"

	"vincent.com/golangginrest/database"
)

func init() {
	database.InitDB()
}

func main() {
	router.InitRouter()
}
