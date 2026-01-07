import { useCallback, useMemo } from 'react'

import { useConfig } from '@/features/settings/states/useConfig'
import type { SymbolInfo } from '@/features/settings/types'

export interface UseSettingsSymbolsReturn {
  availableSymbols: SymbolInfo[]
  selectedSymbols: SymbolInfo[]
  onChange: (symbols: SymbolInfo[]) => void
  formatLabel: (symbol: SymbolInfo) => string
}

export function useSettingsSymbols(): UseSettingsSymbolsReturn {
  const { config, updateConfig, availableSymbols } = useConfig()

  const selectedCoinIds: string[] = config?.symbols ?? []

  const selectedSymbols = useMemo(() => {
    return selectedCoinIds
      .map((coinId) => availableSymbols.find((s) => s.coinId === coinId))
      .filter((s): s is SymbolInfo => !!s)
  }, [availableSymbols, selectedCoinIds])

  const onChange = (symbols: SymbolInfo[]) => {
    updateConfig({ symbols: symbols.map((s) => s.coinId) })
  }

  const formatLabel = useCallback((symbol: SymbolInfo): string => {
    return symbol.symbol === symbol.name
      ? symbol.symbol
      : `${symbol.symbol} - ${symbol.name}`
  }, [])

  return {
    availableSymbols,
    selectedSymbols,
    onChange,
    formatLabel,
  }
}
