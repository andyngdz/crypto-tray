import { CONFIG_DEFAULTS } from '@/features/settings/constants/defaults'
import { useConfig } from '@/features/settings/states/useConfig'

export interface UseSettingsNumberFormatReturn {
  numberFormat: string
  onChange: (value: string) => void
}

export function useSettingsNumberFormat(): UseSettingsNumberFormatReturn {
  const { config, updateConfig } = useConfig()

  const numberFormat = config?.number_format ?? CONFIG_DEFAULTS.numberFormat

  const onChange = (value: string) => {
    if (value) {
      updateConfig({ number_format: value })
    }
  }

  return {
    numberFormat,
    onChange,
  }
}
