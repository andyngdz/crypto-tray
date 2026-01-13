import { FetchPrices, RefreshPrices } from '@bindings/services/appservice'
import { PriceData } from '@bindings/providers/models'

export type { PriceData }

export async function fetchPrices(symbols: string[]): Promise<PriceData[]> {
  const prices = await FetchPrices(symbols)
  return prices.filter((p): p is PriceData => !!p)
}

export function refreshPrices(): void {
  RefreshPrices()
}
