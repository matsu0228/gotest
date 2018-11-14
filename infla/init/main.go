package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// create database
	_, err = db.Exec(`
	CREATE DATABASE IF NOT EXISTS todo
  `)
	if err != nil {
		panic(err)
	}

	// CREATE TABLE
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS todo.tasks(
	  id INT PRIMARY KEY AUTO_INCREMENT, 
	  title VARCHAR(100), 
	  body VARCHAR(1000), 
	  created_at TIMESTAMP NOT NULL default current_timestamp, 
	  updated_at TIMESTAMP NOT NULL default current_timestamp on update current_timestamp 
	)
  `)
	if err != nil {
		panic(err)
	}
}
