import type { Key } from '@heroui/react'
import { useMemo, useState } from 'react'

import type { SymbolInfo } from '@/features/settings/types'

import { useConfig } from '@/features/settings/states/useConfig'

export interface UseSettingsSymbolsReturn {
  inputValue: string
  setInputValue: (value: string) => void
  filteredSymbols: SymbolInfo[]
  selectedSymbols: SymbolInfo[]
  onSelect: (key: Key | null) => void
  onRemove: (keys: Set<Key>) => void
}

export function useSettingsSymbols(): UseSettingsSymbolsReturn {
  const { config, updateConfig, availableSymbols } = useConfig()
  const [inputValue, setInputValue] = useState('')

  const selectedCoinIds: string[] = config?.symbols ?? []

  // Filter out already-selected symbols and apply search filter
  const filteredSymbols = useMemo(() => {
    const search = inputValue.toLowerCase()
    return availableSymbols.filter((s) => {
      if (selectedCoinIds.includes(s.coinId)) return false
      if (!inputValue) return true
      return s.symbol.toLowerCase().includes(search) || s.name.toLowerCase().includes(search)
    })
  }, [availableSymbols, selectedCoinIds, inputValue])

  // Get symbol info for selected values
  const selectedSymbols = useMemo(() => {
    return selectedCoinIds
      .map((coinId) => availableSymbols.find((s) => s.coinId === coinId))
      .filter((s): s is SymbolInfo => !!s)
  }, [availableSymbols, selectedCoinIds])

  const onSelect = (key: Key | null) => {
    if (key) {
      const coinId = key as string
      if (!selectedCoinIds.includes(coinId)) {
        updateConfig({ symbols: [...selectedCoinIds, coinId] })
      }
      setInputValue('')
    }
  }

  const onRemove = (keys: Set<Key>) => {
    const toRemove = new Set([...keys].map((k) => k as string))
    const updated = selectedCoinIds.filter((id) => !toRemove.has(id))
    if (updated.length > 0) {
      updateConfig({ symbols: updated })
    }
  }

  return {
    inputValue,
    setInputValue,
    filteredSymbols,
    selectedSymbols,
    onSelect,
    onRemove,
  }
}
