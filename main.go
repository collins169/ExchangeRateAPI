package main

import (
	database "ExchangeRate/configuration"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	routes "ExchangeRate/routes"
)


func handleRequests(){
	// creates a new instance of a mux router
    router := mux.NewRouter()
    router.HandleFunc("/", routes.Index).Methods(http.MethodGet)
	api := router.PathPrefix("/api/v1").Subrouter()
    api.HandleFunc("/rates", routes.GetAllRates).Methods(http.MethodGet)
    api.HandleFunc("/rates/{code}", routes.GetRatesByCode)
    api.HandleFunc("/rates/convert", routes.ConvertRates).Queries("base", "{base}", "amount", "{amount:[0-999]+}").Methods(http.MethodGet)
    log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	database.IntialDB()
	log.Println("============================================")
	log.Println("Rest API v1.0 - Mux Routers")
	handleRequests()
}