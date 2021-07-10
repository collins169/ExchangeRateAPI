package configuration

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	model "ExchangeRate/model"
)

func IntialDB(){
	os.Remove("sqlite-database.db") // I delete the file to avoid duplicated records. 
	// SQLite is a file based database.

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	database, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer database.Close() // Defer Closing the database
	CreateTable(database)

	log.Println("Let insert the rates records")
	InsertRecord(database, &model.ExchangeRate{0,"NGN", "GHS", 0.014, 69.26})
	InsertRecord(database, &model.ExchangeRate{0,"NGN", "KSH", 0.26, 3.80})
	InsertRecord(database, &model.ExchangeRate{0,"GHS", "NGN", 69.26, 0.014})
	InsertRecord(database, &model.ExchangeRate{0,"GHS", "KSH", 18.20, 0.055})
	InsertRecord(database, &model.ExchangeRate{0,"KSH", "NGN", 3.80, 0.26})
	InsertRecord(database, &model.ExchangeRate{0,"KSH", "GHS", 0.055, 18.20})
}
func Connect() *sql.DB {
	database, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil{
		panic(err)
	}
	return database
}

func Disconnect(db *sql.DB) error{
	return db.Close()
}

func CreateTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE exchange_rates (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"currency_from" TEXT,
		"currency_to" TEXT,
		"conversion_value" TEXT,
		"inverse_conversion_value" NUMERIC
	  );` // SQL Statement for Create Table

	log.Println("Create Exchange Rates table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("exchange rates table created")
}

func InsertRecord(db *sql.DB, rates *model.ExchangeRate){
	log.Println("Inserting rates into database")
	query := `INSERT INTO exchange_rates (currency_from, currency_to, conversion_value, inverse_conversion_value) VALUES (:1, :2, :3, :4)`
	
	_ , err := db.Exec(query, &rates.CurrencyFrom, &rates.CurrencyTo, &rates.ConversionValue, &rates.InverseConversionValue)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func FindAll(db * sql.DB) (*[]model.ExchangeRate, error ){
	rates := []model.ExchangeRate{}
	rows, err := db.Query("SELECT id, currency_from, currency_to, conversion_value, inverse_conversion_value FROM exchange_rates")
	if err != nil{
		log.Fatalln(err.Error())
		return &rates, err
	}
	defer rows.Close()
	for rows.Next(){
		rate := model.ExchangeRate{}
		rows.Scan(&rate.Id, &rate.CurrencyFrom, &rate.CurrencyTo, &rate.ConversionValue, &rate.InverseConversionValue)
		rates = append(rates, rate)
	}
	rows.Close()
	return &rates, nil
}

func FindByCode(db * sql.DB, code string) (*[]model.ExchangeRate, error) {
	rates := []model.ExchangeRate{}
	rows, err := db.Query("SELECT id, currency_from, currency_to, conversion_value, inverse_conversion_value FROM exchange_rates where currency_from = :1", code)
	if err != nil{
		log.Fatalln(err.Error())
		return &rates, err
	}
	defer rows.Close()
	for rows.Next(){
		rate := model.ExchangeRate{}
		rows.Scan(&rate.Id, &rate.CurrencyFrom, &rate.CurrencyTo, &rate.ConversionValue, &rate.InverseConversionValue)
		rates = append(rates, rate)
	}
	rows.Close()
	return &rates, nil
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