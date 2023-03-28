package main

import (
	"log"
	"net/http"

	"gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"

	"github.com/gagassurya19/go-jwt-mux/middlewares"
	"github.com/gagassurya19/go-jwt-mux/controllers/authController"
	"github.com/gagassurya19/go-jwt-mux/controllers/productController"
	"github.com/gagassurya19/go-jwt-mux/models"
)

func main() {
	r := mux.NewRouter()
	models.ConnectDatabase()

	r.HandleFunc("/login", authController.Login).Methods("POST")
	r.HandleFunc("/register", authController.Register).Methods("POST")
	r.HandleFunc("/logout", authController.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/product", productController.Index).Methods("GET")

	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
