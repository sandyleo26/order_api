package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sandyleo26/lalamove/order"
)

func main() {
	uc := order.NewRealUseCase()
	http.Handle("/", Router(uc))
	fmt.Println("Listening :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//Router returns router
func Router(uc order.UseCase) *mux.Router {
	r := mux.NewRouter()
	createOrderHandler := order.NewCreateOrderHandler(uc)
	takeOrderHandler := order.NewTakeOrderHandler(uc)
	getOrderHandler := order.NewGetOrderHandler(uc)
	r.HandleFunc("/order", createOrderHandler).Methods("POST")
	r.HandleFunc("/order/{id}", takeOrderHandler).Methods("PUT")
	r.HandleFunc("/orders", getOrderHandler).Methods("GET")
	return r
}
