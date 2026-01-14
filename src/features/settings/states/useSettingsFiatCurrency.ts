import { CONFIG_DEFAULTS } from '@/features/settings/constants/defaults'
import { useConfig } from '@/features/settings/states/useConfig'

export interface UseSettingsFiatCurrencyReturn {
  displayCurrency: string
  onChange: (value: string) => void
}

export function useSettingsFiatCurrency(): UseSettingsFiatCurrencyReturn {
  const { config, updateConfig } = useConfig()

  const displayCurrency = config?.display_currency ?? CONFIG_DEFAULTS.displayCurrency

  const onChange = (value: string) => {
    if (value) {
      updateConfig({ display_currency: value })
    }
  }

  return {
    displayCurrency,
    onChange,
  }
}
