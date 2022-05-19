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
	r.HandleFunc("/createUser/{name},{surname},{mail},{phone},{id},{pass}", createUser)
	r.HandleFunc("/loginUser/{id},{pass}", loginUser)
	r.HandleFunc("/deleteUser/{id},{pass}", deleteUser)
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

func createUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]
	surname := vars["surname"]
	mail := vars["mail"]
	phone := vars["phone"]
	id := vars["id"]
	pass := vars["pass"]
	isCreate := db.CreateUser(name, surname, mail, phone, id, pass)
	if isCreate {
		w.Write([]byte(name + " hello , account succesfull."))
	} else {
		w.Write([]byte("Account not created. Please check"))
	}

}

func loginUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	pass := vars["pass"]
	isLogin := db.Login(id, pass)

	if isLogin {
		w.Write([]byte("Signed in."))
	} else {
		w.Write([]byte("Failed to login."))
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		vars := mux.Vars(r)
		id := vars["id"]
		pass := vars["pass"]
		isDelete := db.DeleteUser(id, pass)

		if isDelete {
			w.Write([]byte("The account has been successfully deleted."))
		} else {
			w.Write([]byte("The account could not be deleted."))
		}
	} else {
		w.Write([]byte("Only delete method!"))
	}

}
