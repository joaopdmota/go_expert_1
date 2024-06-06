package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type ExchangeRateResponse struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(res)
	}
	if res.StatusCode >= 500 {
		fmt.Println("Error: ", res.Status)
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error: Failed to parse txt file")
		return
	}
	writeExchange(body)
}

func writeExchange(body []byte) error {
	var c ExchangeRateResponse
	err := json.Unmarshal(body, &c)
	if err != nil {
		fmt.Println("Error: Failed to convert data")
	}

	f, err := os.Create("dollar-exchange.txt")

	if err != nil {
		fmt.Println("Error: Failed to create file")
	}

	_, err = f.Write([]byte(fmt.Sprintf("Dólar: %v", c.Bid)))

	if err != nil {
		fmt.Println("Error: failed to write file")
	}

	fmt.Println("File generated successfully!")
	return nil
}
