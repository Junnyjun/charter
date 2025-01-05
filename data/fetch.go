package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type TimeSeriesDaily struct {
	MetaData   map[string]interface{}       `json:"Meta Data"`
	TimeSeries map[string]map[string]string `json:"Time Series (Daily)"`
}

type DailyData struct {
	Date   string
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

func FetchDailyData(symbol, apiKey string) ([]DailyData, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY_ADJUSTED&symbol=%s&outputsize=compact&apikey=%s", symbol, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tsd TimeSeriesDaily
	if err := json.Unmarshal(body, &tsd); err != nil {
		return nil, err
	}

	var data []DailyData
	for date, values := range tsd.TimeSeries {
		open, _ := strconv.ParseFloat(values["1. open"], 64)
		high, _ := strconv.ParseFloat(values["2. high"], 64)
		low, _ := strconv.ParseFloat(values["3. low"], 64)
		close, _ := strconv.ParseFloat(values["4. close"], 64)
		volume, _ := strconv.ParseFloat(values["6. volume"], 64)

		daily := DailyData{
			Date:   date,
			Open:   open,
			High:   high,
			Low:    low,
			Close:  close,
			Volume: volume,
		}
		data = append(data, daily)
	}

	return data, nil
}
