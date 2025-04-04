// services/stock.go
package services

import (
	"math"
	"math/rand"
	"time"
)

// StockService simulates stock price data generation
type StockService struct {
	Price float64
}

// NewStockService initializes a new StockService with a starting price
func NewStockService() *StockService {
	return &StockService{Price: 100.0}
}

// GenerateData simulates the generation of stock price data
// It uses a normal distribution to simulate realistic price movements
func (s *StockService) GenerateData() (time.Time, float64) {
	change := rand.NormFloat64() * 1.5
	// Randomly increase the volatility of the price change
	if rand.Float64() > 0.7 {
		change *= 3
	}
	s.Price += change // Update the price based on the change
	// Ensure the price does not go below a certain threshold
	s.Price = math.Max(s.Price, 50)
	return time.Now().UTC(), math.Round(s.Price*100) / 100
}
