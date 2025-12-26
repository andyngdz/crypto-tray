import { useMemo } from 'react'
import type { Key } from '@heroui/react'
import type { ProviderInfo } from '../types'
import { useConfig } from './useConfig'

export interface UseSettingsProviderReturn {
  providerId: string
  providers: ProviderInfo[]
  currentProvider?: ProviderInfo
  onChange: (value: Key | null) => void
}

export function useSettingsProvider(): UseSettingsProviderReturn {
  const { config, providers, updateConfig } = useConfig()

  const providerId = config?.provider_id ?? ''

  const currentProvider = useMemo(
    () => providers.find((p) => p.id === providerId),
    [providers, providerId]
  )

  const onChange = (value: Key | null) => {
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
