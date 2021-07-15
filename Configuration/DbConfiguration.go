package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	model "ExchangeRate/model"
	// _ "github.com/mattn/go-sqlite3"
)

func IntialDB(){
	os.Remove("rates.json") // I delete the file to avoid duplicated records. 
	// SQLite is a file based database.

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("rates.json") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("rates.json created")
	// CreateTable(database)

	log.Println("Let insert the rates records")
	rates := [] *model.ExchangeRate{}
	rates = append(rates, &model.ExchangeRate{1,"NGN", "GHS", 0.014, 69.26})
	rates = append(rates, &model.ExchangeRate{2,"NGN", "KSH", 0.26, 3.80})
	rates = append(rates, &model.ExchangeRate{3,"GHS", "NGN", 69.26, 0.014})
	rates = append(rates, &model.ExchangeRate{4,"GHS", "KSH", 18.20, 0.055})
	rates = append(rates, &model.ExchangeRate{5,"KSH", "NGN", 3.80, 0.26})
	rates = append(rates, &model.ExchangeRate{6,"KSH", "GHS", 0.055, 18.20})
	data, _ := json.MarshalIndent(rates, "", " ")
	_ = ioutil.WriteFile("rates.json", data, 0644)
}

func FindAll() (*[]model.ExchangeRate, error ){
	rates := []model.ExchangeRate{}
	data, _ := ioutil.ReadFile("rates.json")
	err := json.Unmarshal(data, &rates)
	if err != nil {
		return &rates, err
	}
	log.Println(rates)
	return &rates, nil
}

func FindByCode(code string) (*[]model.ExchangeRate, error) {
	rates := []model.ExchangeRate{}
	rates2 := []model.ExchangeRate{}
	data, err := ioutil.ReadFile("rates.json")
	if err != nil{
		log.Fatalln(err.Error())
		return &rates2, err
	}

	err2 := json.Unmarshal(data, &rates)
	if err2 != nil {
		return &rates2, err2
	}

	for _, rate := range rates {
		if rate.CurrencyFrom == code {
			rates2 = append(rates2, rate)
		}
	}
	return &rates2, nil
}

func IsCurrencyAllowed(code string) bool{
	if code != "" {
		for _, v := range []string{"GHS", "NGN", "KSH"}{
			if v == code{
				return true
			}
		}
	}
	return false
}