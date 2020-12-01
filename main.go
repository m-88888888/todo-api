package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Task struct {
	ID      int64 `gorm:"primaryKey"`
	Name    string
	Remarks string
}

func initDB() (db *gorm.DB) {
	err := godotenv.Load()
	if err != nil {
		// .env読めなかった場合の処理
		log.Fatal("env file loading error")
	}

	DNS := os.Getenv("NAME") + ":" + os.Getenv("PASS") + "@" + os.Getenv("PROTOCOL") + "/" + os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(DNS)

	db, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})

	db.AutoMigrate(&Task{})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("db connection is successful")
	return
}

func getTask(w http.ResponseWriter, r *http.Request) {
	db := initDB()

	id := r.FormValue("id")
	var task Task
	db.First(&task, "id = ?", id)
	fmt.Println(task)
	json.NewEncoder(w).Encode(task)
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	db := initDB()

	var tasks []Task
	db.Find(&tasks)
	fmt.Println(tasks)
	json.NewEncoder(w).Encode(tasks)
}

func insertTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	task.Name = r.FormValue("name")
	task.Remarks = r.FormValue("remarks")
	fmt.Println(task)

	db := initDB()
	result := db.Create(&task)
	if result.RowsAffected > 0 {
		fmt.Println("task insert success")
	} else {
		log.Fatal("Failed to insert task")
	}

}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	var task Task
	task.ID, _ = strconv.ParseInt(id, 10, 64)
	fmt.Println(task)

	db := initDB()
	db.Debug().Delete(&task)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	var task Task
	task.ID, _ = strconv.ParseInt(id, 10, 64)
	task.Name = r.FormValue("name")
	task.Remarks = r.FormValue("remarks")

	db := initDB()
	db.Debug().Model(&task).Where("id = ?", id).Updates(task)
}

func handleRequests() {
	// ハンドラ。URLパスと関数を結びつける
	http.HandleFunc("/task", getTask)
	http.HandleFunc("/tasks", getAllTasks)
	http.HandleFunc("/insert", insertTask)
	http.HandleFunc("/delete", deleteTask)
	http.HandleFunc("/update", updateTask)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
