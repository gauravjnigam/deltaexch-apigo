package main

import (
	"fmt"

	DeltaExchApi "github.com/gauravjnigam/deltaexchapigo"
)

func main() {

	apiKey := "0009omTbCGr2Vo0DKt8cqAf7CW9lpJ"
	apiSecret := "d6U8RiIo9UWfVYaqJSQkjVr3Jkvsh6k6E3j461Xtnqya154a32q1QFwqV06l"
	baseUrl := "https://api.india.delta.exchange"
	// Create New Shoonya Broking Client
	deltaExchClient := DeltaExchApi.New(baseUrl, apiKey, apiSecret, "")

	// assets, err := deltaExchClient.GetAssets()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("Assets :- ", assets)
	start := "22-12-2024 09:15:00"
	end := "25-12-2024 15:10:00"

	resp, err := deltaExchClient.GetTimePriceSeries("BTCUSD", start, end, "15m")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)

}
