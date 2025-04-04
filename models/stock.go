package models

import (
	"encoding/json"
	"time"
)

// StockData represents the structure of stock data sent over WebSocket
type StockData struct {
	Time  time.Time `json:"time"`
	Price float64   `json:"price"`
}

// Custom JSON marshaling to format time
func (s *StockData) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Time  string  `json:"time"`
		Price float64 `json:"price"`
	}{
		Time:  s.Time.Format(time.RFC3339),
		Price: s.Price,
	})
}
