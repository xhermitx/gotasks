package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct{
	Uid int 
	Username string
	PasswordHash string
}

type Task struct{
	TaskName string
	Date string
	Status string
	Uid int
}

func generateHash(pwd string) string{
	h := sha256.New()

	h.Write([]byte(pwd))

	hash := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return hash
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
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Minute*3)

	// CREATING A NEW USER
	if err := createUser(db);err!=nil{
		fmt.Println("Error occured!")
	}

	if err := addTask(db);err!=nil{
		fmt.Println("Error occured")
	}
	
}

func createUser(db *gorm.DB) error{
	var username,password string

	fmt.Println("Enter the username and to be deleted and its password: ")
	fmt.Scan(&username,&password)

	user := User{Username: username, PasswordHash: generateHash(password)}

	result := db.Create(&user)
	if result.Error != nil{
		log.Printf("Error creating user: %v", result.Error)
        return result.Error // Return the error to the caller
	}else{
		log.Print("User ID: ",user)
		log.Print("Rows Affected: ", result.RowsAffected)
	}
	return nil
}

func addTask(db *gorm.DB) error{
	var username,password string

	fmt.Println("Enter the username and to be deleted and its password: ")
	fmt.Scan(&username,&password)

	user := User{Username: username, PasswordHash: generateHash(password)}

	res := db.Find(&user, "username = ? and password_hash = ?",username,generateHash(password))
	if res.Error !=nil{
		log.Print(res.Error)
	}

	// log.Print(user.Uid)

	task := Task{
					TaskName: "Pay bills", 
					Date: time.Now().Format("2006-01-02"), 
					Status: "Pending",
					Uid: user.Uid,
				}

	result := db.Create(&task)

	if result.Error != nil{
		log.Printf("Error creating user: %v", result.Error)
        return result.Error // Return the error to the caller
	}else{
		log.Print("User ID: ",user)
		log.Print("Rows Affected: ", result.RowsAffected)
	}
	return nil
}