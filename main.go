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
	fmt.Println(apiClient.GetQuote().QuoteResponse.Result[0].Ask)

}
