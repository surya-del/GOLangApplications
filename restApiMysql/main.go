package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/orders", createOrder).Methods("POST")

	router.HandleFunc("/orders/{orderId}", getOrder).Methods("GET")

	router.HandleFunc("/orders", getOrders).Methods("GET")

	router.HandleFunc("/orders/{orderId}", updateOrder).Methods("PUT")

	router.HandleFunc("/orders/{orderId}", deleteOrder).Methods("DELETE")

	initDB()

	log.Fatal(http.ListenAndServe(":9000", router))
}

func main() {
	initializeRouter()
}
