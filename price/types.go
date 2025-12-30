package price

import "crypto-tray/providers"

// Callback is called when price data is fetched
type Callback func(data []*providers.PriceData, err error)
