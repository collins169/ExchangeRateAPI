package model

type ExchangeRate struct {
	Id						int	`json:"id"`
	CurrencyFrom			string	`json:"currency_from"`
	CurrencyTo				string	`json:"currency_to"`
	ConversionValue			float64	`json:"conversion_value"`
	InverseConversionValue	float64	`json:"inverse_conversion_value"`
}

type Response struct {
	Message			string	`json:"message"`
	Success			bool	`json:"success"`
	Data			[]ExchangeRate	`json:"data"`
}

type Response1 struct {
	Message			string	`json:"message"`
	Success			bool	`json:"success"`
	Base			string	`json:"based"`
	Data			map[string]float64	`json:"data"`
}