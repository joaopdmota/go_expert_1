package interfaces

import "time"

type ExchangeRateResponseDTO struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type ExchangeRateResponse struct {
	Bid string `json:"bid"`
}

type ExchangeRecord struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Value     string    `json:"value"`
}
