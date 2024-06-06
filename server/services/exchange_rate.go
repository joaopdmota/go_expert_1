package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/database"
	"server/interfaces"
	"time"
)

func FetchExchangeRate() (*interfaces.ExchangeRateResponse, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var c interfaces.ExchangeRateResponseDTO
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return &interfaces.ExchangeRateResponse{
		Bid: c.Usdbrl.Bid,
	}, nil
}

func FetchAllExchangeRate() ([]interfaces.ExchangeRecord, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	result, err := database.GetAllRecords()
	if err != nil {
		return nil, err
	}

	return result, nil
}
