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

func createUser(name,surname,id,pass,mail,phone string)bool{
	fmt.Println("Json reading...")
	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	r := []byte(byteValue)
	data := RouteData{}
	json.Unmarshal(r, &data)
	fmt.Println("Json reading done...")
	fmt.Println("DB connected...")
	var db *sql.DB
	db, err = sql.Open("sqlite3", "./data/likya.db")
	route, err := db.Prepare("INSERT INTO route(routeName) values (?)")
	checkerr.CheckError(err)
	
	res1, err := route.Exec(data.Route.Metadata.Name)
	checkerr.CheckError(err)
	id, err:= res1.LastInsertId()
	checkerr.CheckError(err)
	fmt.Println("Route Name Insert")


	defer func() {
		err = db.Close()
		checkerr.CheckError(err)
		fmt.Println("DB Closed.")
	}()

	return true
}