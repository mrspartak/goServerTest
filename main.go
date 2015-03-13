package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/go/{name:[a-z0-9]+}", profile).Methods("GET")
	rtr.HandleFunc("/go", profile2).Methods("GET")

	http.Handle("/", rtr)
	http.ListenAndServe(":8000", nil)
}

func profile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}
func profile2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
