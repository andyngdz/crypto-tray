export interface Config {
  provider_id: string
  api_keys: Record<string, string>
  refresh_seconds: number
  symbol: string
}

export interface ProviderInfo {
  id: string
  name: string
  requiresApiKey: boolean
}

// Type guards for runtime validation
export function isConfig(obj: unknown): obj is Config {
  if (typeof obj !== 'object' || obj === null) return false
  const c = obj as Record<string, unknown>
  return (
    typeof c.provider_id === 'string' &&
    typeof c.refresh_seconds === 'number' &&
    typeof c.symbol === 'string' &&
    typeof c.api_keys === 'object' &&
    c.api_keys !== null
  )
}

export function isProviderInfo(obj: unknown): obj is ProviderInfo {
  if (typeof obj !== 'object' || obj === null) return false
  const p = obj as Record<string, unknown>
  return (
    typeof p.id === 'string' &&
    typeof p.name === 'string' &&
    typeof p.requiresApiKey === 'boolean'
  )
}
