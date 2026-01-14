import { CONFIG_DEFAULTS } from '@/features/settings/constants/defaults'
import { useConfig } from '@/features/settings/states/useConfig'

export interface UseSettingsAutoStartReturn {
  autoStart: boolean
  onChange: (value: boolean) => void
}

export function useSettingsAutoStart(): UseSettingsAutoStartReturn {
  const { config, updateConfig } = useConfig()

  const autoStart = config?.auto_start ?? CONFIG_DEFAULTS.autoStart

  const onChange = (value: boolean) => {
    updateConfig({ auto_start: value })
  }

  return {
    autoStart,
    onChange,
  }
}
