package main

// go mod init github.com/pjw1702/go-jwt-mux

// go get -u github.com/golang-jwt/jwt/v4
// go get -u github.com/gorilla/mux gorm.io/gorm gorm.io/driver/mysql golang.org/x/crypto

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pjw1702/go-jwt-mux/controllers/authcontroller"
	"github.com/pjw1702/go-jwt-mux/controllers/productcontroller"
	"github.com/pjw1702/go-jwt-mux/middlewares"
	"github.com/pjw1702/go-jwt-mux/models"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
