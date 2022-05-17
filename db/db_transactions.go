package db

import (
	"database/sql"
	"fmt"
	checkerr "likyaapi/checkErr"

	_ "github.com/mattn/go-sqlite3"
)

func DbConnect() {
	fmt.Println("DB connected...")
	var db *sql.DB
	db, err := sql.Open("sqlite3", "./data/likya.db")
	checkerr.CheckError(err)
	tableCreate(db)
}

func tableCreate(db *sql.DB) {

	db.Exec("create table if not exists users (username text, password text)") //Table Create
	fmt.Println("oluşturuldu.")
}
func getId(db *sql.DB){
	db.Query("")
}
