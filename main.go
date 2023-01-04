package main

import (
	"oop/model"
	"oop/nosql"
	"oop/service"
)

func main() {
	user := &model.User{
		Name:     "Edu",
		Password: "1234",
	}

	noSqlUserRepo := nosql.NewNoSqlUserRepository()
	userService := service.NewUserService(noSqlUserRepo)
	userService.InsertUser(user)
}
