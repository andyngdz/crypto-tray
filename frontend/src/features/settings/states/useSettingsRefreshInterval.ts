import { useConfig } from '@/features/settings/states/useConfig'
import type { Key } from '@heroui/react'

export interface UseSettingsRefreshIntervalReturn {
  refreshSeconds: number
  onChange: (value: Key | null) => void
}

export function useSettingsRefreshInterval(): UseSettingsRefreshIntervalReturn {
  const { config, updateConfig } = useConfig()

  const refreshSeconds = config?.refresh_seconds ?? 0

  const onChange = (value: Key | null) => {
    if (value) {
      updateConfig({ refresh_seconds: Number(value) })
    }
  }

  return {
    refreshSeconds,
    onChange,
  }
}
