package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
} 

func ApiMain(){
	r := mux.NewRouter()
	r.HandleFunc("/",Index)

	http.ListenAndServe(":7070",r)
}
