package main

import (
	"example/go-orm-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/go_orm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	// Create
	db.Create(&model.User{Fname: "Karn", Lname: "Yong", Username: "karn.yong@melivecode.com", Avatar: "https://www.melivecode.com/users/1.png"})
	db.Create(&model.User{Fname: "Ivy", Lname: "Cal", Username: "ivy.cal@melivecode.com", Avatar: "https://www.melivecode.com/users/2.png"})

}
