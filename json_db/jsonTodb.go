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

//Route mark information from outside with RouteId adds it to the database.
func insertMark(db *sql.DB, data RouteData, routeId int64) { //RouteId comes from the route table.
	routeMark, err := db.Prepare("INSERT INTO route_mark(routeId, markName, markLat, markLon) values (?,?,?,?)")
	checkerr.CheckError(err)
	for _, k := range data.Route.RouteMark {
		_, err = routeMark.Exec(routeId, k.MarkName, k.MarkLat, k.MarkLon)
		checkerr.CheckError(err)
	}
}

//Route coordinate information from outside with RouteId adds it to the database.
func insertCoordinate(db *sql.DB, data RouteData, routeId int64) {
	routeMark, err := db.Prepare("INSERT INTO route_coordinate(lat, lon, routeId) values (?,?,?)")
	checkerr.CheckError(err)
	for _, k := range data.Route.RouteCoordinate.Coordinates.CoordinateArr {
		_, err = routeMark.Exec(k.Lat, k.Lon, routeId)
		checkerr.CheckError(err)
	}
}

//The main function reads the Json document and converts it to a struct.
//Then it adds it to the database with the appropriate functions.
func JsonToDB(path string) {
	fmt.Println("Json reading...")
	jsonFile, err := os.Open(path)
	checkerr.CheckError(err)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	r := []byte(byteValue)
	data := RouteData{} //Data data turned into RouteData struct.
	json.Unmarshal(r, &data)
	fmt.Println("Json reading done...")
	fmt.Println("DB connected...")
	var db *sql.DB
	db, err = sql.Open("sqlite3", "./data/likya.db")
	route, err := db.Prepare("INSERT INTO route(routeName) values (?)")
	checkerr.CheckError(err)

	res1, err := route.Exec(data.Route.Metadata.Name) //route name insert
	checkerr.CheckError(err)
	id, err := res1.LastInsertId()
	checkerr.CheckError(err)
	fmt.Println("Route Name Insert")

	insertMark(db, data, id)
	fmt.Println("Route Mark Insert")
	insertCoordinate(db, data, id)
	fmt.Println("Route Coordinate Insert")

	//After the file and database process is finished, they close it last.
	defer func() {
		err := jsonFile.Close()
		checkerr.CheckError(err)
		fmt.Println("JsonFile Closed.")
		err = db.Close()
		checkerr.CheckError(err)
		fmt.Println("DB Closed.")
	}()
}
