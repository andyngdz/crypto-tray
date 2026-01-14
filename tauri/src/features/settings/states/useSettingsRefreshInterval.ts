import { CONFIG_DEFAULTS } from '@/features/settings/constants/defaults'
import { useConfig } from '@/features/settings/states/useConfig'

export interface UseSettingsRefreshIntervalReturn {
  refreshSeconds: number
  onChange: (value: string) => void
}

export function useSettingsRefreshInterval(): UseSettingsRefreshIntervalReturn {
  const { config, updateConfig } = useConfig()

  const refreshSeconds = config?.refresh_seconds ?? CONFIG_DEFAULTS.refreshSeconds

  const onChange = (value: string) => {
    if (value) {
      updateConfig({ refresh_seconds: Number(value) })
    }
  }

  return {
    refreshSeconds,
    onChange,
  }
}
