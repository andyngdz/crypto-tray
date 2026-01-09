package movement

import (
	"sync"

	"crypto-tray/providers"
)

// Tracker tracks price movements between fetches
type Tracker struct {
	mu         sync.RWMutex
	lastPrices map[string]float64
	movements  map[string]Direction
}

// NewTracker creates a new movement tracker
func NewTracker() *Tracker {
	return &Tracker{
		lastPrices: make(map[string]float64),
		movements:  make(map[string]Direction),
	}
}

// Track compares current prices to last known and returns movement directions
func (t *Tracker) Track(data []*providers.PriceData) map[string]Direction {
	t.mu.Lock()
	defer t.mu.Unlock()

	result := make(map[string]Direction, len(data))

	for dataIdx := range data {
		d := data[dataIdx]
		currentPrice := d.ConvertedPrice
		if currentPrice == 0 {
			currentPrice = d.Price
		}

		var direction Direction
		if lastPrice, exists := t.lastPrices[d.CoinID]; exists {
			if currentPrice > lastPrice {
				direction = Up
			} else if currentPrice < lastPrice {
				direction = Down
			} else {
				// Price unchanged - keep previous direction
				direction = t.movements[d.CoinID]
			}
		} else {
			direction = Neutral // First fetch
		}

		t.lastPrices[d.CoinID] = currentPrice
		t.movements[d.CoinID] = direction
		result[d.CoinID] = direction
	}

	return result
}
