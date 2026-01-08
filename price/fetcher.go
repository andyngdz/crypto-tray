package price

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"crypto-tray/config"
	"crypto-tray/providers"
)

// Fetcher periodically fetches cryptocurrency prices
type Fetcher struct {
	registry      *providers.Registry
	configManager *config.Manager
	callback      Callback

	mu         sync.Mutex
	cancelFunc context.CancelFunc
	refreshCh  chan struct{}
}

// NewFetcher creates a new price fetcher
func NewFetcher(
	registry *providers.Registry,
	configManager *config.Manager,
	callback Callback,
) *Fetcher {
	return &Fetcher{
		registry:      registry,
		configManager: configManager,
		callback:      callback,
		refreshCh:     make(chan struct{}, 1),
	}
}

// Start begins the price fetching loop
func (f *Fetcher) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	f.mu.Lock()
	f.cancelFunc = cancel
	f.mu.Unlock()

	go f.loop(ctx)
}

// Stop stops the price fetching loop
func (f *Fetcher) Stop() {
	f.mu.Lock()
	if f.cancelFunc != nil {
		f.cancelFunc()
	}
	f.mu.Unlock()
}

// RefreshNow triggers an immediate price fetch
func (f *Fetcher) RefreshNow() {
	select {
	case f.refreshCh <- struct{}{}:
	default:
		// Already a refresh pending
	}
}

func (f *Fetcher) loop(ctx context.Context) {
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
			f.fetchOnce(ctx, cfg)
			// Update ticker if interval changed
			if cfg.RefreshSeconds != currentInterval {
				currentInterval = cfg.RefreshSeconds
				ticker.Reset(time.Duration(currentInterval) * time.Second)
			}
		case <-f.refreshCh:
			cfg = f.configManager.Get()
			f.fetchOnce(ctx, cfg)
			// Reset ticker after manual refresh
			currentInterval = cfg.RefreshSeconds
			ticker.Reset(time.Duration(currentInterval) * time.Second)
		}
	}
}

func (f *Fetcher) fetchOnce(ctx context.Context, cfg config.Config) {
	provider, ok := f.registry.Get(cfg.ProviderID)
	if !ok {
		log.Printf("Unknown provider: %s", cfg.ProviderID)
		f.callback(nil, fmt.Errorf("unknown provider: %s", cfg.ProviderID))
		return
	}

	if len(cfg.Symbols) == 0 {
		f.callback([]*providers.PriceData{}, nil)
		return
	}

	data, err := provider.FetchPrices(ctx, cfg.Symbols)
	f.callback(data, err)
}
