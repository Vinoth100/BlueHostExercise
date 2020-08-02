package server

import (
	// "fmt"
	"log"
	"net/http"

	"acme/pkg/config"
	"acme/pkg/product"

	"github.com/gorilla/mux"
)

//Start and Routing the API
func Start(conf config.Config) {
	myRouter := mux.NewRouter().StrictSlash(true)
	log.Printf(conf.ServerPort)
	myRouter.HandleFunc("/product", product.LoadProduct).Methods("POST")
	myRouter.HandleFunc("/product/add", product.AddProduct).Methods("PUT")
	myRouter.HandleFunc("/product", product.ListProducts).Methods("GET")
	myRouter.HandleFunc("/product/email", product.EmailSchdedule).Methods("GET")
	log.Fatal(http.ListenAndServe(conf.ServerPort, myRouter))
}
