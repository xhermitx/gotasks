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
	Username string
	PasswordHash string
}

type Task struct{
	TaskName string
	TaskDate time.Time
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
	user := User{Username: "Rohit", PasswordHash: generateHash("pwd")}

	result := db.Create(&user)
	if result.Error != nil{
		log.Fatal(result.Error)
	}else{
		log.Print("User ID: ",user)
		log.Print("Rows Affected: ",result.RowsAffected)
	}
}