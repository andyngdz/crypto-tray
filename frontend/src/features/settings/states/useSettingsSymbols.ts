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

  const selectedIds: string[] = config?.symbols ?? []

  // Filter out already-selected symbols and apply search filter
  const filteredSymbols = useMemo(() => {
    return availableSymbols.filter((s) => {
      // Exclude already selected
      if (selectedIds.includes(s.id)) return false
      // Apply search filter (case-insensitive on id and name)
      if (!inputValue) return true
      const search = inputValue.toLowerCase()
      return s.id.toLowerCase().includes(search) || s.name.toLowerCase().includes(search)
    })
  }, [availableSymbols, selectedIds, inputValue])

  // Get symbol info for selected values
  const selectedSymbols = useMemo(() => {
    return selectedIds
      .map((id) => availableSymbols.find((s) => s.id === id))
      .filter((s): s is SymbolInfo => !!s)
  }, [availableSymbols, selectedIds])

  const onSelect = (key: Key | null) => {
    if (key) {
      const symbol = key as string
      if (!selectedIds.includes(symbol)) {
        updateConfig({ symbols: [...selectedIds, symbol] })
      }
      setInputValue('')
    }
  }

  const onRemove = (keys: Set<Key>) => {
    const toRemove = new Set([...keys].map((k) => k as string))
    const updated = selectedIds.filter((s) => !toRemove.has(s))
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
