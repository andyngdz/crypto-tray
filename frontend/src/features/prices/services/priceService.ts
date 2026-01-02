import type { PriceData } from '@/features/prices/types'
import { isPriceData } from '@/features/prices/types'
import { FetchPrices, RefreshPrices } from '@wailsjs/go/main/App'

export async function fetchPrices(symbols: string[]): Promise<PriceData[]> {
  const prices = await FetchPrices(symbols)
  if (!Array.isArray(prices) || !prices.every(isPriceData)) {
    throw new Error('Invalid prices response from backend')
  }
  return prices
}

export function refreshPrices(): void {
  RefreshPrices()
}
