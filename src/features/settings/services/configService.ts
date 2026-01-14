import { invoke } from '@tauri-apps/api/core'
import { getCurrentWindow } from '@tauri-apps/api/window'
import type {
  Config,
  ProviderInfo,
  SymbolInfo,
} from '@/features/settings/types'
import {
  isConfig,
  isProviderInfo,
  isSymbolInfo,
} from '@/features/settings/types'

/**
 * Fetches the current configuration from the Rust backend
 */
export async function fetchConfig(): Promise<Config> {
  const cfg = await invoke<Config>('get_config')

  if (!isConfig(cfg)) {
    throw new Error('Invalid config response from backend')
  }

  return cfg
}

/**
 * Fetches the list of available API providers from the Rust backend
 */
export async function fetchProviders(): Promise<ProviderInfo[]> {
  const provs = await invoke<ProviderInfo[]>('get_available_providers')

  if (!Array.isArray(provs) || !provs.every(isProviderInfo)) {
    throw new Error('Invalid providers response from backend')
  }

  return provs
}

/**
 * Saves the configuration to the Rust backend
 */
export async function saveConfig(config: Config): Promise<void> {
  await invoke('save_config', { config })
}

/**
 * Fetches the list of available cryptocurrency symbols
 */
export async function fetchAvailableSymbols(): Promise<SymbolInfo[]> {
  const symbols = await invoke<SymbolInfo[]>('get_available_symbols')

  if (!Array.isArray(symbols) || !symbols.every(isSymbolInfo)) {
    throw new Error('Invalid symbols response from backend')
  }

  return symbols
}

/**
 * Hides the settings window
 */
export async function hideWindow(): Promise<void> {
  const window = getCurrentWindow()
  await window.hide()
}
