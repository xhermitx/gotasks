package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	models "github.com/xhermitx/gotasks/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the home Page")
	fmt.Println("Endpoint hit: homepage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/tasks",handleViewTasks)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func handleViewTasks(w http.ResponseWriter, r *http.Request){
	fmt.Println("EndPoint Hit: handleViewTasks")
	json.NewEncoder(w).Encode(models.Task{Tid: 1,TaskName: "test",Status: "test",Date: "test"})
}

func main() {

	handleRequests()

	err := godotenv.Load()
	if err!=nil{
		log.Panic("Error loading the environment variables")
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",dbUser,dbPassword,dbAddress,dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!= nil{
		log.Panic(err)
	}

	sqlDB, err := db.DB()

	// SET THE MAXIMUM NUMBER OF CONNECTIONS AND OPEN CONNECTIONS
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)

	// SET THE MAXIMUM TIME FOR A CONNECTION TO BE REUSED
	sqlDB.SetConnMaxLifetime(time.Minute*3)

}

