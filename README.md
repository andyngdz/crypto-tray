# CryptoTray

A lightweight cryptocurrency price tracker that lives in your system tray.

## Features

- Real-time cryptocurrency price updates in system tray
- Support for multiple crypto providers (CoinGecko, Binance)
- Customizable refresh intervals
- Multiple fiat currency support
- Start on system startup (enabled by default)
- Cross-platform: Windows, macOS, Linux

## Screenshots

Coming soon.

## Requirements

- Go 1.24+
- Node.js 20+
- pnpm 9+
- Wails v2.11+

### Windows (CGO requirement)

This project uses `github.com/emersion/go-autostart` which requires CGO on Windows.

1. Install MSYS2 from https://www.msys2.org/
2. Open "MSYS2 MINGW64" terminal and run:
   ```bash
   pacman -S mingw-w64-x86_64-gcc
   ```
3. Add to your system PATH: `C:\msys64\mingw64\bin`
4. Set environment variable: `CGO_ENABLED=1`

### macOS

Xcode command line tools (includes clang):
```bash
xcode-select --install
```

### Linux

```bash
sudo apt-get install build-essential libgtk-3-dev libwebkit2gtk-4.1-dev
```

## Development

```bash
# Install Wails CLI
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# Install frontend dependencies
cd frontend && pnpm install && cd ..

# Run in development mode
wails dev
```

## Building

```bash
wails build
```

## License

[MIT License](LICENSE)
