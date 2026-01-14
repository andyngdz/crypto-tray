import { invoke } from '@tauri-apps/api/core'

export interface PriceData {
  coin_id: string
  symbol: string
  name: string
  price: number
  change_24h: number
  movement?: string
}

export async function fetchPrices(symbols: string[]): Promise<PriceData[]> {
  const prices = await invoke<PriceData[]>('fetch_prices', { symbols })
  return prices.filter((p): p is PriceData => !!p)
}

export function refreshPrices(): void {
  invoke('refresh_prices')
}
