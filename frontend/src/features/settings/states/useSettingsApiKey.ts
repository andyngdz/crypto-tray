import { useConfig } from './useConfig'

export interface UseSettingsApiKeyReturn {
  value: string
  onChange: (value: string) => void
}

export function useSettingsApiKey(): UseSettingsApiKeyReturn {
  const { config, updateConfig } = useConfig()

  const value = config?.api_keys[config.provider_id] ?? ''

  const onChange = (newValue: string) => {
    if (!config) return
    updateConfig({
      api_keys: {
        ...config.api_keys,
        [config.provider_id]: newValue,
      },
    })
  }

  return {
    value,
    onChange,
  }
}
