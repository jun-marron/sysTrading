package main

import (
	"fmt"
	"log"
	"sysTrading/config"
	"sysTrading/utils"
	"sysTrading/yahooFinance"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)
	log.Println("test")
	apiClient := yahooFinance.New(config.Config.ApiKey)
	//quote, _ := apiClient.GetQuote("SPCE")
	//fmt.Println(quote.QuoteResponse.Result[0])
	chart, _ := apiClient.GetChart("", "1y", "US", "1d", "en", "div,split", "SPXL")
	fmt.Println(chart)
}
