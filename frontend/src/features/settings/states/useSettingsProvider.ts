import { useConfig } from '@/features/settings/states/useConfig'
import type { ProviderInfo } from '@/features/settings/types'
import type { Key } from '@heroui/react'
import { useMemo } from 'react'

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
