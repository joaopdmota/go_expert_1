package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/database"
	"server/services"
)

func FetchExchangeRateHandler(w http.ResponseWriter, r *http.Request) {
	response, err := services.FetchExchangeRate()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = database.StoreData(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Request processed")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func FetchAllExchangeRateHandler(w http.ResponseWriter, r *http.Request) {
	response, err := services.FetchAllExchangeRate()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Request processed")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
