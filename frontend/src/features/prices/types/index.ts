export interface PriceData {
  symbol: string
  price: number
  change_24h: number
}

export function isPriceData(obj: unknown): obj is PriceData {
  if (typeof obj !== 'object' || obj === null) return false
  const p = obj as Record<string, unknown>
  return (
    typeof p.symbol === 'string' &&
    typeof p.price === 'number' &&
    typeof p.change_24h === 'number'
  )
}
