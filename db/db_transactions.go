package db

import (
	"database/sql"
	"fmt"
	checkerr "likyaapi/checkErr"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name     string
	Surname  string
	Mail     string
	Phone    int
	userName string
	Password string
}
type RouteCoordinate struct {
	RouteName string
	Lat       float32
	Long      float32
}

var routeCoordinateArr []RouteCoordinate
var lastString string

func GetRoute() {
	routeCoordinateArr = nil
	fmt.Println("DB connected...")
	db, err := sql.Open("sqlite3", "./data/likya.db")
	checkerr.CheckError(err)

	rows, err := db.Query("Select * from route")
	checkerr.CheckError(err)
	for rows.Next() {
		var routeName string
		err = rows.Scan(&routeName)
		checkerr.CheckError(err)
		lastString = routeName
	}
	defer db.Close()
}

// func tableCreate(db *sql.DB) {

// 	db.Exec("create table if not exists users (username text, password text)") //Table Create
// 	fmt.Println("oluşturuldu.")
// }
