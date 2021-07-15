package main

import (
	database "ExchangeRate/configuration"
	"log"
	"net/http"
	routes "ExchangeRate/routes"
)


func handleRequests(){
	// creates a new instance of a mux router
    http.HandleFunc("/api/v1/", routes.Index)
    http.HandleFunc("/api/v1/rates", routes.GetAllRates)
    http.HandleFunc("/api/v1/rates/convert", routes.ConvertRates)
    http.HandleFunc("/api/v1/rates/", routes.GetRatesByCode)
    log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	database.IntialDB()
	log.Println("============================================")
	log.Println("Rest API v1.0 - Mux Routers")
	handleRequests()
}