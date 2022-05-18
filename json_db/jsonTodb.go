package jsonTodb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	checkerr "likyaapi/checkErr"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type RouteData struct {
	Route struct {
		Metadata struct {
			Name string `json:"name"`
		} `json:"metadata"`
		RouteMark []struct {
			MarkLat  string `json:"-lat"`
			MarkLon  string `json:"-lon"`
			MarkName string `json:"name"`
		} `json:"wpt"`
		RouteCoordinate struct {
			Name        string `json:"name"`
			Coordinates struct {
				CoordinateArr []struct {
					Lat string `json:"-lat"`
					Lon string `json:"-lon"`
				} `json:"trkpt"`
			} `json:"trkseg"`
		} `json:"trk"`
	} `json:"gpx"`
}
func insertMark(db *sql.DB,data RouteData,routeId int64){
	routeMark, err := db.Prepare("INSERT INTO route_mark(routeId, markName, markLat, markLon) values (?,?,?,?)")
	checkerr.CheckError(err)
	for _ ,k := range data.Route.RouteMark{
		_, err =routeMark.Exec(routeId,k.MarkName,k.MarkLat,k.MarkLon)
		checkerr.CheckError(err)
	}
}
func insertCoordinate(db *sql.DB,data RouteData,routeId int64){
	routeMark, err := db.Prepare("INSERT INTO route_coordinate(lat, lon, routeId) values (?,?,?)")
	checkerr.CheckError(err)
	for _ ,k := range data.Route.RouteCoordinate.Coordinates.CoordinateArr{
		_, err =routeMark.Exec(k.Lat,k.Lon,routeId)
		checkerr.CheckError(err)
	}
}
func JsonToDB(path string) {
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

	insertMark(db,data,id)
	fmt.Println("Route Mark Insert")
	insertCoordinate(db,data,id)
	fmt.Println("Route Coordinate Insert")


	defer func() {
		err := jsonFile.Close()
		checkerr.CheckError(err)
		fmt.Println("JsonFile Closed.")
		err = db.Close()
		checkerr.CheckError(err)
		fmt.Println("DB Closed.")
	}()
}
