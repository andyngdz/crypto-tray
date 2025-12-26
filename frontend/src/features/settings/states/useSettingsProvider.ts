import { useMemo } from 'react'
import type { SharedSelection } from '@heroui/react'
import type { ProviderInfo } from '../types'
import { useConfig } from './useConfig'

export interface UseSettingsProviderReturn {
  providerId: string
  providers: ProviderInfo[]
  currentProvider?: ProviderInfo
  onChange: (keys: SharedSelection) => void
}

export function useSettingsProvider(): UseSettingsProviderReturn {
  const { config, providers, updateConfig } = useConfig()

  const providerId = config?.provider_id ?? ''

  const currentProvider = useMemo(
    () => providers.find((p) => p.id === providerId),
    [providers, providerId]
  )

  const onChange = (keys: SharedSelection) => {
    if (keys === 'all') return
    const value = Array.from(keys)[0]
    if (value) {
      updateConfig({ provider_id: String(value) })
    }
  }

  return {
    providerId,
    providers,
    currentProvider,
    onChange,
  }
}
