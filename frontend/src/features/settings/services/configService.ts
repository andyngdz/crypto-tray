import type { Config, ProviderInfo, SymbolInfo } from '@/features/settings/types'
import { isConfig, isProviderInfo, isSymbolInfo } from '@/features/settings/types'
import {
  GetAvailableProviders,
  GetAvailableSymbols,
  GetConfig,
  HideWindow,
  SaveConfig,
} from '@wailsjs/go/main/App'

/**
 * Fetches the current configuration from the Go backend
 */
export async function fetchConfig(): Promise<Config> {
  const cfg = await GetConfig()
  if (!isConfig(cfg)) {
    throw new Error('Invalid config response from backend')
  }
  return cfg
}

/**
 * Fetches the list of available API providers from the Go backend
 */
export async function fetchProviders(): Promise<ProviderInfo[]> {
  const provs = await GetAvailableProviders()
  if (!Array.isArray(provs) || !provs.every(isProviderInfo)) {
    throw new Error('Invalid providers response from backend')
  }
  return provs
}

/**
 * Saves the configuration to the Go backend
 */
export async function saveConfig(config: Config): Promise<void> {
  await SaveConfig(config)
}

/**
 * Fetches the list of available cryptocurrency symbols
 */
export async function fetchAvailableSymbols(): Promise<SymbolInfo[]> {
  const symbols = await GetAvailableSymbols()
  if (!Array.isArray(symbols) || !symbols.every(isSymbolInfo)) {
    throw new Error('Invalid symbols response from backend')
  }
  return symbols
}

/**
 * Hides the settings window
 */
export function hideWindow(): void {
  HideWindow()
}
