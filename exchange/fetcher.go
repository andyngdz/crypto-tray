package exchange

import (
	"context"
	"sync"
	"time"

	"crypto-tray/config"
	"crypto-tray/internal/httpclient"
)

// Fetcher periodically fetches exchange rates
type Fetcher struct {
	configManager  *config.Manager
	callback       RatesHandler
	primaryClient  *httpclient.Client
	fallbackClient *httpclient.Client

	mu           sync.RWMutex
	cancelFunc   context.CancelFunc
	refreshCh    chan struct{}
	currentRates *ExchangeRates
}

// newFetcher creates a new exchange rate fetcher (internal use by Service)
func newFetcher(configManager *config.Manager) *Fetcher {
	return &Fetcher{
		configManager: configManager,
		primaryClient: httpclient.New(httpclient.Config{
			BaseURL: primaryURL,
			Timeout: timeout,
		}),
		fallbackClient: httpclient.New(httpclient.Config{
			BaseURL: fallbackURL,
			Timeout: timeout,
		}),
		refreshCh: make(chan struct{}, 1),
	}
}

// Start begins the exchange rate fetching loop with the given callback
func (f *Fetcher) Start(callback RatesHandler) {
	f.mu.Lock()
	f.callback = callback
	f.mu.Unlock()

	ctx, cancel := context.WithCancel(context.Background())

	f.mu.Lock()
	f.cancelFunc = cancel
	f.mu.Unlock()

	go f.loop(ctx)
}

// Stop stops the exchange rate fetching loop
func (f *Fetcher) Stop() {
	f.mu.Lock()
	if f.cancelFunc != nil {
		f.cancelFunc()
	}
	f.mu.Unlock()
}

// RefreshNow triggers an immediate exchange rate fetch
func (f *Fetcher) RefreshNow() {
	select {
	case f.refreshCh <- struct{}{}:
	default:
		// Already a refresh pending
	}
}

// GetRates returns the current cached exchange rates (thread-safe)
func (f *Fetcher) GetRates() *ExchangeRates {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.currentRates
}

func (f *Fetcher) loop(ctx context.Context) {
	// Fetch immediately on start
	f.fetchOnce(ctx)

	cfg := f.configManager.Get()
	currentInterval := cfg.RefreshSeconds
	ticker := time.NewTicker(time.Duration(currentInterval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			cfg = f.configManager.Get()
			f.fetchOnce(ctx)
			// Update ticker if interval changed
			if cfg.RefreshSeconds != currentInterval {
				currentInterval = cfg.RefreshSeconds
				ticker.Reset(time.Duration(currentInterval) * time.Second)
			}
		case <-f.refreshCh:
			cfg = f.configManager.Get()
			f.fetchOnce(ctx)
			// Reset ticker after manual refresh
			currentInterval = cfg.RefreshSeconds
			ticker.Reset(time.Duration(currentInterval) * time.Second)
		}
	}
}

func (f *Fetcher) fetchOnce(ctx context.Context) {
	rates, err := f.fetchFromClient(ctx, f.primaryClient)
	if err != nil {
		rates, err = f.fetchFromClient(ctx, f.fallbackClient)
		if err != nil {
			// Keep using cached rates
			f.callback(f.GetRates(), nil)
			return
		}
	}

	// Update cache with successful fetch
	f.mu.Lock()
	f.currentRates = rates
	f.mu.Unlock()

	f.callback(rates, nil)
}

func (f *Fetcher) fetchFromClient(ctx context.Context, client *httpclient.Client) (*ExchangeRates, error) {
	var resp APIResponse

	if err := client.Get(ctx, "/"+baseCurrency+".json", &resp); err != nil {
		return nil, err
	}

	return &ExchangeRates{
		Date:  resp.Date,
		Base:  baseCurrency,
		Rates: resp.USDT,
	}, nil
}
