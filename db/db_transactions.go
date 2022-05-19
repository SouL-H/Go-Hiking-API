package db

import (
	"database/sql"

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
	Lat       float64
	Lon       float64
}

var routeCoordinateArr []RouteCoordinate

func GetRoute(routeId int) []RouteCoordinate {
	routeCoordinateArr = nil

	db, err := sql.Open("sqlite3", "./data/likya.db")
	checkerr.CheckError(err)

	rows, err := db.Query("SELECT route.routeName, route.routeId,route_coordinate.lat,  route_coordinate.lon FROM route_coordinate INNER JOIN route ON route.routeId = route_coordinate.routeId WHERE route.routeId=?",routeId)
	checkerr.CheckError(err)
	for rows.Next() {
		var lat float64
		var lon float64
		var routeId int64
		var routeName string
		err := rows.Scan(&routeName, &routeId, &lon, &lat)
		if err == nil {
			routeCoordinateArr = append(routeCoordinateArr, RouteCoordinate{
				RouteName: routeName,
				Lat:       lat,
				Lon:       lon,
			})
		}

	}
	db.Close()
	return routeCoordinateArr
	
	
}

// func tableCreate(db *sql.DB) {

// 	db.Exec("create table if not exists users (username text, password text)") //Table Create
// 	fmt.Println("oluşturuldu.")
// }
