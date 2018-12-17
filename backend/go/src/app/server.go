package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/jobs", handler).Methods("GET")
	http.ListenAndServe(":30002", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{"jobs": []string{"job1", "job2"}})
}
