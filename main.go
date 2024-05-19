package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	handlers "github.com/xhermitx/gotasks/handlers"
	msql "github.com/xhermitx/gotasks/store/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DEFINE THE HOME PAGE
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home Page")
	fmt.Println("Endpoint hit: homepage")
}

// HANDLE THE ROUTES
func handleRequests(handler *handlers.TaskHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/tasks", handler.ViewTasks).Methods("GET")
	router.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", handler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(os.Getenv("ADDRESS"), router))
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading the environment variables")
	}

	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	mysqlDB := msql.NewMySQLStore(db)
	taskHandler := handlers.NewTaskHandler(mysqlDB)

	handleRequests(taskHandler)
}
