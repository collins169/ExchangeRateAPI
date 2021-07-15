package routes

import (
	database "ExchangeRate/configuration"
	model "ExchangeRate/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
    w.Write([]byte("Api Works!\n"))
}

func GetAllRates(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	log.Println("About to fetch all Rates")
	response := model.Response{}
	rates, err := database.FindAll()
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		response.Data = *rates
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Message = "Success"
	response.Success = true
	response.Data = *rates
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

func GetRatesByCode(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	log.Println("About to fetch all Rates by currency code")
    code := r.URL.Query().Get("code")
	response := model.Response1{}
	if !database.IsCurrencyAllowed(code){
		response.Message = fmt.Sprintf("Currency %s is not allowed", code)
		response.Success = false
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(response)
		return
	}
	rates, err := database.FindByCode(code)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Message = "Success"
	response.Success = true
	response.Base = code
	data := make(map[string]float64)
	for _, rate := range *rates{
		data[rate.CurrencyTo] = rate.ConversionValue
	}
	response.Data = data
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

func ConvertRates(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	log.Println("About to fetch all Rates by currency code")
    code := r.URL.Query().Get("base")
    amount := r.URL.Query().Get("amount")
	response := model.Response1{}
	if !database.IsCurrencyAllowed(code){
		response.Message = fmt.Sprintf("Currency %s is not allowed", code)
		response.Success = false
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(response)
		return
	}
	rates, err := database.FindByCode(code)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return
	}
	data := make(map[string]float64)
	if amt, err := strconv.ParseFloat(amount, 64); err == nil {
		for _, rate := range *rates{
			data[rate.CurrencyTo] = rate.ConversionValue * amt
		}
	}
	response.Message = "Success"
	response.Success = true
	response.Base = code
	response.Data = data
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}