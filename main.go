package main

import (
	"echo-practice/controller"
	"echo-practice/db"
	"echo-practice/repository"
	"echo-practice/router"
	"echo-practice/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
