package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Shoe struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}

var db *gorm.DB
var err error

func main() {
	// connet db by giving alll instances using gorm.Open
	// tak a r instance for mux
	// define all routes
	// start server

	dsn := "host=localhost user=dhiran pasword=pass123 dbname=shoes_store port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Shoe{})

	r := mux.NewRouter()

	r.HandleFunc("/shoes", getShoes).Methods("GET")
	r.HandleFunc("/shoe", getShoe).Methods("GET")
	r.HandleFunc("/shoe", createShoe).Methods("POST")
	r.HandleFunc("/shoe", updateShoe).Methods("PUT")
	r.HandleFunc("/shoe", delete).Methods("DELETE")

	fmt.Println("Server starting at localhost:8080 port")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getShoes(rw http.ResponseWriter, r *http.Request) {
	var shoes []Shoe
	db.Find(&shoes)
	fmt.Println("All shoes retrieved successfully")
}

func getShoe(rw http.ResponseWriter, r *http.Request) {
	var shoe Shoe
	params := mux.Vars(r)
	db.First(&shoe, params["id"])
	// json.NewDecoder(r.Body).Decode(&shoe)
	// db.Save(&shoe)
}

func createShoe(rw http.ResponseWriter, r *http.Request) {
	var shoe Shoe
	json.NewDecoder(r.Body).Decode(&shoe)
	db.Create(&shoe)
}

func updateShoe(rw http.ResponseWriter, r *http.Request) {
	var shoe Shoe
	json.NewDecoder(r.Body).Decode(&shoe)
	db.Save(&shoe)
}

func delete(rw http.ResponseWriter, r *http.Request) {
	var shoe Shoe
	param := mux.Vars(r)

	db.First(&shoe, param["id"])
	db.Delete(&shoe)
	fmt.Println("Shoe details deleted successfully")
}
