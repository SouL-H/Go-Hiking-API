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
	RouteName     string
	CoordinateArr []Coordinate
}

type Coordinate struct {
	Lat float64
	Lon float64
}

var routeCoordinate RouteCoordinate
var routeName string
var coordinateArr []Coordinate

func GetRoute(routeId int) RouteCoordinate {
	coordinateArr = nil

	db, err := sql.Open("sqlite3", "./data/likya.db")
	checkerr.CheckError(err)

	rows, err := db.Query("SELECT route.routeName, route.routeId,route_coordinate.lat,  route_coordinate.lon FROM route_coordinate INNER JOIN route ON route.routeId = route_coordinate.routeId WHERE route.routeId=?", routeId)
	checkerr.CheckError(err)
	for rows.Next() {
		var lat float64
		var lon float64
		var routeId int64
		var _routeName string
		err := rows.Scan(&_routeName, &routeId, &lon, &lat)
		if err == nil {
			coordinateArr = append(coordinateArr, Coordinate{
				Lat: lat,
				Lon: lon,
			})
			routeName = _routeName
		}

	}
	routeCoordinate.RouteName = routeName
	routeCoordinate.CoordinateArr = coordinateArr
	db.Close()
	return routeCoordinate

}

// func tableCreate(db *sql.DB) {

// 	db.Exec("create table if not exists users (username text, password text)") //Table Create
// 	fmt.Println("oluşturuldu.")
// }
