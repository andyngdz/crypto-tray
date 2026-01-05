import { useMemo } from 'react'

import { useConfig } from '@/features/settings/states/useConfig'
import type { ProviderInfo } from '@/features/settings/types'

export interface UseSettingsProviderReturn {
  providerId: string
  providers: ProviderInfo[]
  currentProvider?: ProviderInfo
  onChange: (value: string) => void
}

export function useSettingsProvider(): UseSettingsProviderReturn {
  const { config, providers, updateConfig } = useConfig()

  const providerId = config?.provider_id ?? ''

  const currentProvider = useMemo(
    () => providers.find((p) => p.id === providerId),
    [providers, providerId]
  )

  const onChange = (value: string) => {
    if (value) {
      updateConfig({ provider_id: value })
    }
  }

  return {
    providerId,
    providers,
    currentProvider,
    onChange,
  }
}
