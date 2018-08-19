package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/order", CreateOrder).Methods("POST")
	r.HandleFunc("/order/{id}", TakeOrder).Methods("PUT")
	r.HandleFunc("/orders", GetOrder).Methods("GET")
	http.Handle("/", r)
	fmt.Println("Listening :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create Order")
}

func TakeOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Take order")
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get Order")
}
