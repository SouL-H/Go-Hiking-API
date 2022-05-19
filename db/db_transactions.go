package db

import (
	"database/sql"
	"fmt"

	checkerr "likyaapi/checkErr"
	"likyaapi/crypto"

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

type MarkCoordinate struct {
	MarkLat  float64
	MarkLon  float64
	MarkName string
}

type RouteMark struct {
	RouteName    string
	RouteMarkArr []MarkCoordinate
}

var routeCoordinate RouteCoordinate
var routeName string
var coordinateArr []Coordinate

var routeMark RouteMark
var markCoordinateArr []MarkCoordinate
//Returns coordinates based on route id.
func GetRoute(routeId int) RouteCoordinate {
	coordinateArr = nil
	routeName = ""
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
	rows.Close()

	return routeCoordinate

}
//Returns the coordinate and name of the marks on the route.
func GetRouteMark(routeId int) RouteMark {
	markCoordinateArr = nil
	routeName = ""
	db, err := sql.Open("sqlite3", "./data/likya.db")
	checkerr.CheckError(err)

	rows, err := db.Query("SELECT  route.routeName, route.routeId,route_mark.markName, route_mark.markLat,  route_mark.markLon  FROM route_mark INNER JOIN route ON route.routeId = route_mark.routeId WHERE route.routeId=?", routeId)
	checkerr.CheckError(err)
	for rows.Next() {
		var markLat float64
		var markLon float64
		var routeId int64
		var _routeName string
		var markName string
		err := rows.Scan(&_routeName, &routeId, &markName, &markLat, &markLon)
		if err == nil {
			markCoordinateArr = append(markCoordinateArr, MarkCoordinate{
				MarkName: markName,
				MarkLat:  markLat,
				MarkLon:  markLon,
			})
			routeName = _routeName
		}

	}
	routeMark.RouteName = routeName
	routeMark.RouteMarkArr = markCoordinateArr
	db.Close()
	rows.Close()

	return routeMark

}
//A new user is created.
func CreateUser(name, surname, mail, phone, id, pass string) bool {
	var db *sql.DB
	db, err := sql.Open("sqlite3", "./data/likya.db")
	route, err := db.Prepare("INSERT INTO users(name,surname,mail,phone,userName,userPass) values (?,?,?,?,?,?)")
	checkerr.CheckError(err)
	hashPass, _ := crypto.HashPassword(pass)

	_, err = route.Exec(name, surname, mail, phone, id, hashPass)
	checkerr.CheckError(err)

	defer func() {
		err = db.Close()
		checkerr.CheckError(err)
		fmt.Println("DB Closed.")
	}()

	return true
}
//Deletes the current user.
func DeleteUser(id, pass string) bool {
	db, err := sql.Open("sqlite3", "./data/likya.db")
	checkerr.CheckError(err)

	if err == nil {

		rows, err := db.Query("SELECT * FROM users WHERE  userName=?", id)
		for rows.Next() {
			var name, surname, mail, phone, id, _pass string
			var tid int

			err := rows.Scan(&tid, &name, &surname, &mail, &phone, &id, &_pass)
			if err == nil {
				rows.Close()
				hash := crypto.CheckPasswordHash(pass, _pass)
				if hash {
					del, err := db.Prepare("DELETE FROM users where userName=?")
					checkerr.CheckError(err)

					_, err = del.Exec(id)
					checkerr.CheckError(err)
					return true
				}
			}

		}
		checkerr.CheckError(err)
		return false
	}

	db.Close()
	return false
}
//Login the system.
func Login(id, pass string) bool {
	db, err := sql.Open("sqlite3", "./data/likya.db")
	checkerr.CheckError(err)

	rows, err := db.Query("SELECT * FROM users WHERE  userName=?", id)
	for rows.Next() {
		var name, surname, mail, phone, id, _pass string
		var tid int

		err := rows.Scan(&tid, &name, &surname, &mail, &phone, &id, &_pass)
		if err == nil {
			hash := crypto.CheckPasswordHash(pass, _pass)
			if hash {
				return true
			}
		}

	}
	checkerr.CheckError(err)

	defer func() {
		db.Close()
		rows.Close()
	}()

	return false
}
