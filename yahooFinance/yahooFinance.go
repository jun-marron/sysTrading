package yahooFinance

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sysTrading/config"
)

const baseURL = "https://yfapi.net"

type APIClient struct {
	Key        string
	httpClient *http.Client
}

func New(key string) *APIClient {
	apiCleint := &APIClient{key, &http.Client{}}
	return apiCleint
}

func (api *APIClient) doRequest(method, urlPath string, query map[string]string, data []byte) (body []byte, err error) {
	baseURL, err := url.Parse(baseURL)
	if err != nil {
		log.Printf("action=doRequest err=%s", err.Error())
		return nil, err
	}

	apiURL, err := url.Parse(urlPath)
	if err != nil {
		log.Printf("action=doRequest err=%s", err.Error())
		return nil, err
	}

	endpoint := baseURL.ResolveReference(apiURL).String()
	log.Printf("action=doRequest endpoint=%s", endpoint)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("action=doRequest err=%s", err.Error())
		return nil, err
	}
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	req.Header.Add("x-api-key", config.Config.ApiKey)
	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}

type Quote struct {
	QuoteResponse struct {
		Error  interface{} `json:"error"`
		Result []struct {
			Ask                               float64 `json:"ask"`
			AskSize                           int     `json:"askSize"`
			AverageDailyVolume10Day           int     `json:"averageDailyVolume10Day"`
			AverageDailyVolume3Month          int     `json:"averageDailyVolume3Month"`
			Bid                               float64 `json:"bid"`
			BidSize                           int     `json:"bidSize"`
			BookValue                         float64 `json:"bookValue"`
			Currency                          string  `json:"currency"`
			DisplayName                       string  `json:"displayName"`
			DividendDate                      int     `json:"dividendDate"`
			EarningsTimestamp                 int     `json:"earningsTimestamp"`
			EarningsTimestampEnd              int     `json:"earningsTimestampEnd"`
			EarningsTimestampStart            int     `json:"earningsTimestampStart"`
			EpsCurrentYear                    float64 `json:"epsCurrentYear"`
			EpsForward                        float64 `json:"epsForward"`
			EpsTrailingTwelveMonths           float64 `json:"epsTrailingTwelveMonths"`
			EsgPopulated                      bool    `json:"esgPopulated"`
			Exchange                          string  `json:"exchange"`
			ExchangeDataDelayedBy             int     `json:"exchangeDataDelayedBy"`
			ExchangeTimezoneName              string  `json:"exchangeTimezoneName"`
			ExchangeTimezoneShortName         string  `json:"exchangeTimezoneShortName"`
			FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
			FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
			FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
			FiftyTwoWeekHigh                  float64 `json:"fiftyTwoWeekHigh"`
			FiftyTwoWeekHighChange            float64 `json:"fiftyTwoWeekHighChange"`
			FiftyTwoWeekHighChangePercent     float64 `json:"fiftyTwoWeekHighChangePercent"`
			FiftyTwoWeekLow                   float64 `json:"fiftyTwoWeekLow"`
			FiftyTwoWeekLowChange             float64 `json:"fiftyTwoWeekLowChange"`
			FiftyTwoWeekLowChangePercent      float64 `json:"fiftyTwoWeekLowChangePercent"`
			FiftyTwoWeekRange                 string  `json:"fiftyTwoWeekRange"`
			FinancialCurrency                 string  `json:"financialCurrency"`
			FirstTradeDateMilliseconds        int64   `json:"firstTradeDateMilliseconds"`
			ForwardPE                         float64 `json:"forwardPE"`
			FullExchangeName                  string  `json:"fullExchangeName"`
			GmtOffSetMilliseconds             int     `json:"gmtOffSetMilliseconds"`
			Language                          string  `json:"language"`
			LongName                          string  `json:"longName"`
			Market                            string  `json:"market"`
			MarketCap                         int64   `json:"marketCap"`
			MarketState                       string  `json:"marketState"`
			MessageBoardID                    string  `json:"messageBoardId"`
			PostMarketChange                  float64 `json:"postMarketChange"`
			PostMarketChangePercent           float64 `json:"postMarketChangePercent"`
			PostMarketPrice                   float64 `json:"postMarketPrice"`
			PostMarketTime                    int     `json:"postMarketTime"`
			PriceEpsCurrentYear               float64 `json:"priceEpsCurrentYear"`
			PriceHint                         int     `json:"priceHint"`
			PriceToBook                       float64 `json:"priceToBook"`
			QuoteSourceName                   string  `json:"quoteSourceName"`
			QuoteType                         string  `json:"quoteType"`
			Region                            string  `json:"region"`
			RegularMarketChange               float64 `json:"regularMarketChange"`
			RegularMarketChangePercent        float64 `json:"regularMarketChangePercent"`
			RegularMarketDayHigh              float64 `json:"regularMarketDayHigh"`
			RegularMarketDayLow               float64 `json:"regularMarketDayLow"`
			RegularMarketDayRange             string  `json:"regularMarketDayRange"`
			RegularMarketOpen                 float64 `json:"regularMarketOpen"`
			RegularMarketPreviousClose        float64 `json:"regularMarketPreviousClose"`
			RegularMarketPrice                float64 `json:"regularMarketPrice"`
			RegularMarketTime                 int     `json:"regularMarketTime"`
			RegularMarketVolume               int     `json:"regularMarketVolume"`
			SharesOutstanding                 int64   `json:"sharesOutstanding"`
			ShortName                         string  `json:"shortName"`
			SourceInterval                    int     `json:"sourceInterval"`
			Symbol                            string  `json:"symbol"`
			Tradeable                         bool    `json:"tradeable"`
			TrailingAnnualDividendRate        float64 `json:"trailingAnnualDividendRate"`
			TrailingAnnualDividendYield       float64 `json:"trailingAnnualDividendYield"`
			TrailingPE                        float64 `json:"trailingPE"`
			Triggerable                       bool    `json:"triggerable"`
			TwoHundredDayAverage              float64 `json:"twoHundredDayAverage"`
			TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
			TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`
		} `json:"result"`
	} `json:"quoteResponse"`
}

func (api *APIClient) GetQuote(symbols string) (Quote, error) {
	var quote Quote
	url := "/v6/finance/quote"
	query := map[string]string{
		"symbols": symbols,
	}
	resp, err := api.doRequest("GET", url, query, nil)
	if err != nil {
		log.Printf("action=GetQuery err=%s", err.Error())
		return quote, err
	}

	err = json.Unmarshal(resp, &quote)
	if err != nil {
		log.Printf("action=GetQuery err=%s", err.Error())
		return quote, err
	}
	return quote, nil
}

type Chart struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Currency             string  `json:"currency"`
				Symbol               string  `json:"symbol"`
				ExchangeName         string  `json:"exchangeName"`
				InstrumentType       string  `json:"instrumentType"`
				FirstTradeDate       int     `json:"firstTradeDate"`
				RegularMarketTime    int     `json:"regularMarketTime"`
				Gmtoffset            int     `json:"gmtoffset"`
				Timezone             string  `json:"timezone"`
				ExchangeTimezoneName string  `json:"exchangeTimezoneName"`
				RegularMarketPrice   float64 `json:"regularMarketPrice"`
				ChartPreviousClose   float64 `json:"chartPreviousClose"`
				PriceHint            int     `json:"priceHint"`
				CurrentTradingPeriod struct {
					Pre struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"pre"`
					Regular struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"regular"`
					Post struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"post"`
				} `json:"currentTradingPeriod"`
				DataGranularity string   `json:"dataGranularity"`
				Range           string   `json:"range"`
				ValidRanges     []string `json:"validRanges"`
			} `json:"meta"`
			Timestamp []int `json:"timestamp"`
			Events    struct {
				Dividends struct {
					Num1624368600 struct {
						Amount float64 `json:"amount"`
						Date   int     `json:"date"`
					} `json:"1624368600"`
					Num1640097000 struct {
						Amount float64 `json:"amount"`
						Date   int     `json:"date"`
					} `json:"1640097000"`
				} `json:"dividends"`
			} `json:"events"`
			Indicators struct {
				Quote []struct {
					Volume []int     `json:"volume"`
					High   []float64 `json:"high"`
					Low    []float64 `json:"low"`
					Close  []float64 `json:"close"`
					Open   []float64 `json:"open"`
				} `json:"quote"`
				Adjclose []struct {
					Adjclose []float64 `json:"adjclose"`
				} `json:"adjclose"`
			} `json:"indicators"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"chart"`
}

func (api *APIClient) GetChart(comparisons, period, region, interval, lang, events, ticker string) (Chart, error) {
	var chart Chart
	url := "/v8/finance/chart/" + ticker + "/"
	query := map[string]string{
		"comparisons": comparisons,
		"range":       period,
		"region":      region,
		"interval":    interval,
		"lang":        lang,
		"events":      events,
		"ticker":      ticker,
	}
	resp, err := api.doRequest("GET", url, query, nil)
	if err != nil {
		log.Printf("action=GetQuery err=%s", err.Error())
		return chart, err
	}

	err = json.Unmarshal(resp, &chart)
	if err != nil {
		log.Printf("action=GetQuery err=%s", err.Error())
		return chart, err
	}
	return chart, nil
}
