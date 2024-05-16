package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Task struct{
	Tid      int    `gorm:"primary_key;AUTO_INCREMENT"` // Specify auto-increment
    TaskName string `gorm:"not null"`
    Status   string `gorm:"not null;default:'Pending'"`
    Date     string `gorm:"not null"`
}

func main() {

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

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(10)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Minute*3)

	// ADD A TASK
	// if err := addTask(db);err!=nil{
	// 	fmt.Println("Error occured")
	// }

	// DELETE A TASK
	// if err := deleteTask(db);err!=nil{
	// 	fmt.Println("Error occured")
	// }

	// // UPDATE A TASK
	// if err := updateTask(db);err!=nil{
	// 	fmt.Println("Error occured")
	// }

	// VIEW ALL TASKS
	if err := viewTasks(db);err!=nil{
		fmt.Println("Error occured")
	}

}

func addTask(db *gorm.DB) error{

	task := Task{
					TaskName: "Pay Bills", 
					Date: time.Now().Format("2006-01-02 15:04:05"), 
					Status: "Pending",
				}

	result := db.Create(&task)
	if result.Error != nil{
		log.Printf("Error creating user: %v", result.Error)
        return result.Error // Return the error to the caller
	}
	
	log.Print("Task ID: ",task.Tid)
	log.Print("Rows Affected: ", result.RowsAffected)

	return nil
}

func deleteTask(db *gorm.DB) error{
	var taskID int
	fmt.Println("\nEnter the ID of the task to be deleted: ")
	fmt.Scan(&taskID)

	result := db.Delete(&Task{},taskID)
	if result.Error != nil{
		log.Println("An Error occured while deleting the Task")
		return result.Error
	}

	log.Println("Number of Rows Affected : ",result.RowsAffected)

	return nil
}

func updateTask(db *gorm.DB) error{

	reader := bufio.NewReader(os.Stdin)

	var taskID int
	var taskName, taskDate, taskStatus string

	fmt.Print("Enther the Task ID to be updated : ")
	fmt.Scanln(&taskID)

	fmt.Print("Enter the Task Name to be Updated: ")
	taskName,err := reader.ReadString('\n')
	if err!=nil{
		log.Print("Error occured while scanning taskName")
		return err
	}
	taskName = strings.TrimSpace(taskName)

	fmt.Print("Enter the Task Status to be Updated (Completed, Pending, In Progress): ")
	taskStatus,_ = reader.ReadString('\n')
	taskStatus = strings.TrimSpace(taskStatus)

	fmt.Print("Enter the Task Date to be Updated (YYYY-MM-DD HH-MM-SS): ")
	taskDate,_ = reader.ReadString('\n')
	taskDate = strings.TrimSpace(taskDate)

	result := db.Save(&Task{Tid: taskID,TaskName: taskName,Status: taskStatus,Date: taskDate})
	if result.Error != nil{
		log.Println("Error updating the task")
		return result.Error
	}

	log.Println("Rows Affected : ", result.RowsAffected)
	return nil
}

func viewTasks(db *gorm.DB) error{
	
	var tasks []Task
	
	result := db.Find(&tasks)
	if result.Error !=nil{
		log.Print("Error retrieving Data")
		return result.Error
	}

	log.Println("Rows retrieved: ")
	for _, task := range tasks{
		fmt.Printf("ID: %d, Name: %s, Date: %s, Status: %s\n", task.Tid, task.TaskName, task.Date, task.Status)
	}

	return nil
}