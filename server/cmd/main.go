package main

import (
	"fmt"
	"net/http"
	"server/database"
	handler "server/handler"
)

func main() {
	database.SetupDatabase()
	fmt.Println("Serving on port 8080")
	http.HandleFunc("/cotacao", handler.FetchExchangeRateHandler)
	http.HandleFunc("/cotacao/list", handler.FetchAllExchangeRateHandler)
	http.ListenAndServe(":8080", nil)
}
