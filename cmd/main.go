package main

import (
	"log"
	"net/http"

	"github.com/akhilesh-ingle-ge/crud-web-app/config"
	"github.com/akhilesh-ingle-ge/crud-web-app/controllers"
	"github.com/gorilla/mux"
)

func main() {
	config.SetUpDB()

	r := mux.NewRouter()

	r.HandleFunc("/students", controllers.GetStudents).Methods("GET")
	r.HandleFunc("/students", controllers.CreateStudent).Methods("POST")
	r.HandleFunc("/students/{id}", controllers.DeleteStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
