package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/akhilesh-ingle-ge/crud-web-app/config"
	"github.com/akhilesh-ingle-ge/crud-web-app/models"
)

func GetStudents(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/index.tmpl"))
	// if err != nil {
	// 	// panic("Could not load template..")
	// 	// fmt.Println(fmt.Errorf(err.Error()))
	// 	fmt.Printf("Error is: %v\n", err)
	// }
	DB := config.DB
	var Table []models.Student
	DB.Find(&Table)
	fmt.Println(Table)
	tmpl.Execute(w, Table)
	fmt.Println("Hello")
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
}
