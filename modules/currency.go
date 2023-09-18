package modules

import (
	"encoding/json"
	"log"
	"os"
)

type CurrencyData struct {
	Data map[string]CurrencyRate `json:"currencies"`
}

type CurrencyRate map[string]float64

var Currency CurrencyData

// Read currency infomation file
func LoadCurrencyConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	decode := json.NewDecoder(file)
	if err := decode.Decode(&Currency); err != nil {
		log.Println("Error decoding rate.json:", err)
		return err
	}
	return nil
}
