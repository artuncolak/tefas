package tefas

import "time"

type Fund struct {
	Date            string  `json:"TARIH"`
	Code            string  `json:"FONKODU"`
	Name            string  `json:"FONUNVAN"`
	Price           float64 `json:"FIYAT"`
	UnitCount       float64 `json:"TEDPAYSAYISI"`
	InvestorCount   float64 `json:"KISISAYISI"`
	PortfolioAmount float64 `json:"PORTFOYBUYUKLUK"`
	ExchangePrice   string  `json:"BORSABULTENFIYAT"`
}

type FundType string

const (
	YAT FundType = "YAT"
	EMK FundType = "EMK"
	BYF FundType = "BYF"
)

type FundInfoRequest struct {
	Type      FundType
	Code      string
	StartDate time.Time
	EndDate   time.Time
}

type FundInfoResponse struct {
	Draw            int    `json:"draw"`
	RecordsTotal    int    `json:"recordsTotal"`
	RecordsFiltered int    `json:"recordsFiltered"`
	Data            []Fund `json:"data"`
}
