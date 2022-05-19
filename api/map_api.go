package api

import (
	"encoding/json"
	"fmt"
	checkerr "likyaapi/checkErr"
	"likyaapi/db"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
}

func ApiMain() {
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/getInfoRoute/{id}", getInfoRoute)
	r.HandleFunc("/getInfoRouteMark/{id}", getInfoRouteMark)

	http.ListenAndServe(":7070", r)
}

func getInfoRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	Cint, _ := strconv.Atoi(id)
	b := db.GetRoute(Cint)
	j, err := json.Marshal(b)
	checkerr.CheckError(err)
	fmt.Fprintf(w, string(j))
}

func getInfoRouteMark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	Cint, _ := strconv.Atoi(id)
	b := db.GetRouteMark(Cint)
	j, err := json.Marshal(b)
	checkerr.CheckError(err)
	fmt.Fprintf(w, string(j))
}
