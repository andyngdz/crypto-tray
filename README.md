# CryptoTray

A lightweight cryptocurrency price tracker that lives in your system tray.

## Features

- Real-time cryptocurrency price updates in system tray
- Support for multiple crypto providers
- Customizable refresh intervals
- Multiple fiat currency support
- Start on system startup
- Cross-platform: Windows, macOS, Linux

## Requirements

- Node.js 20+
- pnpm 9+
- Rust toolchain (stable)

### Linux system deps

On Ubuntu/Debian:

```bash
sudo apt-get install -y \
  libgtk-3-dev \
  libwebkit2gtk-4.1-dev \
  libayatana-appindicator3-dev \
  librsvg2-dev \
  build-essential \
  pkg-config
```

## Development

```bash
pnpm install
pnpm tauri dev
```

For quick UI iteration (frontend only):

```bash
pnpm dev
```

## Building

```bash
pnpm tauri build
```

## License

[MIT License](LICENSE)
