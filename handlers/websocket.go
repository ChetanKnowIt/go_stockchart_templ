package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"stockchart/models"
	"stockchart/services"
	"time"

	"github.com/gorilla/websocket"
)

// HandleWebSocket handles WebSocket connections and sends stock data
// to the client at regular intervals (5 seconds as per Ticker)
var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true }, // Allow all origins
	}
	tickerInterval = 5 * time.Second // Default value
)

// Export this function to change interval from main.go
func SetTickerInterval(d time.Duration) {
	tickerInterval = d
}

// Add proper error handling
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	defer conn.Close()

	stockService := services.NewStockService()
	ticker := time.NewTicker(tickerInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			t, price := stockService.GenerateData()
			data := models.StockData{
				Time:  t,
				Price: price,
			}

			jsonData, err := json.Marshal(data)
			if err != nil {
				log.Printf("JSON marshal error: %v", err)
				continue
			}

			if err := conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
				log.Printf("WebSocket write error: %v", err)
				return
			}
		}
	}
}
