package main

import (
	"echo-practice/db"
	"echo-practice/model"
	"fmt"
)

func main() { //$env:GO_ENV="dev"; go run .\migrate\migtate.go

	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
