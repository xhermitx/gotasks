package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_user     = "root"
	db_password = "mysql@1SI18CS096"
	db_address  = "localhost"
	db_name     = "GOTASKS"
)

func main() {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",db_user,db_password,db_address,db_name))
	if err!= nil{
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err!=nil{
		log.Printf("Error connecting to DB: %s",db_name)
	}else{
		log.Print("Successfully connected to DB!")
	}

}