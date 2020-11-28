package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name    string
	Remarks string
}

func initDB() (db *gorm.DB) {
	var err error
	DNS := os.Getenv("NAME") + ":" + os.Getenv("PASS") + "@" + os.Getenv("PROTOCOL") + "/" + os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("db connection is successful")
	return
}

func returnTask() (task Task) {
	err := godotenv.Load()
	if err != nil {
		// .env読めなかった場合の処理
		log.Fatal("env file loading error")
	}

	db := initDB()

	// db.AutoMigrate(&Task{})

	// task := Task{
	// 	Name:    "evans試す",
	// 	Remarks: "https://github.com/ktr0731/evans",
	// }

	// db.Create(&task)

	db.First(&task, 1)
	fmt.Println(task)
	return
}

func main() {
	returnTask()
}
