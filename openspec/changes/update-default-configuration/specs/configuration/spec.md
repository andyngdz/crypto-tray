## ADDED Requirements

### Requirement: Provider Default Symbols
Each provider SHALL define its own list of default cryptocurrency symbols in the provider's native format.

#### Scenario: Binance default symbols
- **WHEN** a new user starts the application with Binance as the default provider
- **THEN** the default selected symbols SHALL be BTCUSDT, ETHUSDT, and SOLUSDT

#### Scenario: CoinGecko default symbols
- **WHEN** a user switches to the CoinGecko provider
- **THEN** the default selected symbols SHALL be bitcoin, ethereum, and solana

#### Scenario: Provider switch resets symbols
- **WHEN** a user changes the provider setting
- **THEN** the selected symbols SHALL be reset to the new provider's default symbols

## MODIFIED Requirements

### Requirement: Default Provider Configuration
The application SHALL use Binance as the default price data provider for new installations.

#### Scenario: New installation defaults
- **WHEN** a user installs the application for the first time
- **THEN** the default provider SHALL be Binance
- **AND** the default fiat currency SHALL be USD
- **AND** the default number format SHALL be US (1,234.56)
- **AND** the default selected cryptocurrencies SHALL be BTC, ETH, and SOL
