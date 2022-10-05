package main

import (
	"fmt"
	"learnGo/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/db_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connection to database is good")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "test simpan dari service"
	userInput.Email = "Contoh@gmail.com"
	userInput.Occupation = "anak Bank"
	userInput.Password = "1234"

	userService.RegisterUser(userInput)

}
