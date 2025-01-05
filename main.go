package charter

import (
	"fmt"
	"log"
)

func main() {
	apiKey := "YOUR_ALPHA_VANTAGE_API_KEY"
	symbol := "AAPL"

	dailyData, err := data.FetchDailyData(symbol, apiKey)
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range dailyData {
		fmt.Printf("날짜: %s, 종가: %.2f, 거래량: %.0f\n", d.Date, d.Close, d.Volume)
	}
}
