package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello, World!")

	_ = connection()
	fmt.Println("db connection is successful")
}

func connection() (db *gorm.DB) {
	dsn := "root:@/tasks?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return
}
