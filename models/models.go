package models

type Student struct {
	SId     int    `gorm:"primaryKey; autoIncrement:true" json:"sid"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Contact string `json:"contact"`
}
