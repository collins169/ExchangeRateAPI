# Go Currency Exchange Rate REST API
A simple golang api for currency conversion

## Installation & Run
```bash
# Download this project
go get github.com/collins169/ExchangeRateAPI
```

```bash
# Build and Run
cd ExchangeRateAPI
go build
./ExchangeRateAPI

# API Endpoint : http://127.0.0.1:3000/api/v1
```

## Structure
```
├── Configuration // contains the database layer
│   └── DbConfiguration.go          // Database and queries for our application
├── Model 
│   └── index.go     // Models for our application      
├── Routes // Our API core handlers
│   └── rates.go     // APIs for Rates
└── main.go
```

## API

#### /rates
* `GET` : Get all rates

#### /rates/:code
* `GET` : Get a rate

#### /rates/convert?base={base}&amount={amount}
* `Get` : Convert amount to all available rates