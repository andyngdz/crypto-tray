import type { Key } from '@heroui/react'

import { useConfig } from '@/features/settings/states/useConfig'

export interface UseSettingsSymbolsReturn {
  value: Key[]
  onChange: (keys: Key[]) => void
}

export function useSettingsSymbols(): UseSettingsSymbolsReturn {
  const { config, updateConfig } = useConfig()

  const value: Key[] = config?.symbols ?? []

  const onChange = (keys: Key[]) => {
    const symbols = keys as string[]
    if (symbols.length > 0) {
      updateConfig({ symbols })
    }
  }

  return {
    value,
    onChange,
  }
}
