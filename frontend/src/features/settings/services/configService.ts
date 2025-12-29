import type { Config, ProviderInfo } from '@/features/settings/types'
import { isConfig, isProviderInfo } from '@/features/settings/types'
import {
  GetAvailableProviders,
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
 * Hides the settings window
 */
export function hideWindow(): void {
  HideWindow()
}
