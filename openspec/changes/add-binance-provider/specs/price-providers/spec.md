## ADDED Requirements

### Requirement: Binance Provider

The system SHALL support Binance as a cryptocurrency price data provider.

#### Scenario: Fetch available symbols
- **WHEN** the user selects Binance as provider
- **THEN** the system fetches USDT trading pairs from Binance exchangeInfo API
- **AND** filters to only TRADING status symbols
- **AND** caches results for 24 hours

#### Scenario: Fetch prices
- **WHEN** prices are requested for selected symbols
- **THEN** the system fetches from Binance ticker/24hr endpoint
- **AND** returns price and 24-hour percentage change for each symbol

#### Scenario: Symbol format
- **WHEN** displaying Binance symbols
- **THEN** coinID is the trading pair (e.g., "BTCUSDT")
- **AND** symbol is the base asset (e.g., "BTC")
- **AND** name is the full coin name from hardcoded map or falls back to symbol

### Requirement: Provider Default Coin

Each provider SHALL specify a default coin ID for its platform.

#### Scenario: CoinGecko default
- **WHEN** CoinGecko is the provider
- **THEN** the default coin ID is "bitcoin"

#### Scenario: Binance default
- **WHEN** Binance is the provider
- **THEN** the default coin ID is "BTCUSDT"

### Requirement: Provider Switch Behavior

The system SHALL reset symbol selection when the user switches providers.

#### Scenario: Switch to different provider
- **WHEN** the user changes from one provider to another
- **THEN** the symbol selection is reset to the new provider's default coin
- **AND** the configuration is saved with the new provider and default symbol
