import type { SharedSelection } from '@heroui/react'
import { useConfig } from './useConfig'

export interface UseSettingsRefreshIntervalReturn {
  refreshSeconds: number
  onChange: (keys: SharedSelection) => void
}

export function useSettingsRefreshInterval(): UseSettingsRefreshIntervalReturn {
  const { config, updateConfig } = useConfig()

  const refreshSeconds = config?.refresh_seconds ?? 0

  const onChange = (keys: SharedSelection) => {
    if (keys === 'all') return
    const value = Array.from(keys)[0]
    if (value) {
      updateConfig({ refresh_seconds: Number(value) })
    }
  }

  return {
    refreshSeconds,
    onChange,
  }
}
