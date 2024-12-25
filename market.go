package deltaexchapigo

import (
	"fmt"
	"net/http"
	"net/url"
)

type AssetsResponse struct {
	ID                  int    `json:"id"`
	Symbol              string `json:"symbol"`
	Precision           int    `json:"precision"`
	DepositStatus       string `json:"deposit_status"`
	WithdrawalStatus    string `json:"withdrawal_status"`
	BaseWithdrawalFee   string `json:"base_withdrawal_fee"`
	MinWithdrawalAmount string `json:"min_withdrawal_amount"`
}

type TSPResponse struct {
	Success bool     `json:"success"`
	Candles []Candle `json:"result"`
}

type Candle struct {
	Time   int `json:"time"`
	Open   int `json:"open"`
	High   int `json:"high"`
	Low    int `json:"low"`
	Close  int `json:"close"`
	Volume int `json:"volume"`
}

type TSPriceParam struct {
	Resolution string `json:"resolution"`
	Symbol     string `json:"symbol"`
	Start      string `json:"start"`
	End        string `json:"end"`
}

// GetLTP gets Last Traded Price.
func (c *Client) GetAssets() (AssetsResponse, error) {
	var assetsResp AssetsResponse
	err := c.doEnvelope(http.MethodGet, URIAssets, nil, nil, &assetsResp, true)
	return assetsResp, err
}

// Get historial timePrice series
func (c *Client) GetTimePriceSeries(symbol string, startTime string, endTime string, interval string) (TSPResponse, error) {
	start := GetTime(startTime)
	end := GetTime(endTime)
	tsPriceParam := TSPriceParam{Resolution: interval, Symbol: symbol, Start: start, End: end}
	var candle []Candle

	fmt.Printf("TSP Param: \n%v\n", tsPriceParam)
	params := structToMap(tsPriceParam, "json")
	fmt.Printf("Req Param: \n%v\n", params)

	urlParams := url.Values{}
	for key, value := range params {

		urlParams.Add(key, value.(string))
	}
	fullURL := URIHistoryCandle + "?" + urlParams.Encode()

	fmt.Printf("FullUrl - %v", fullURL)
	err := c.doEnvelope(http.MethodGet, fullURL, nil, nil, &candle, true)

	var tsResponse TSPResponse
	if err != nil {
		return tsResponse, nil
	}

	tsResponse = TSPResponse{Success: true, Candles: candle}
	return tsResponse, nil
}
