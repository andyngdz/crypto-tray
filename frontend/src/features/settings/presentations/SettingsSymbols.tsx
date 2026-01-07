import { Autocomplete, TextField } from '@mui/material'

import { useSettingsSymbols } from '@/features/settings/states/useSettingsSymbols'
import type { SymbolInfo } from '@/features/settings/types'

export function SettingsSymbols() {
  const { availableSymbols, selectedSymbols, onChange } = useSettingsSymbols()

  return (
    <Autocomplete
      multiple
      options={availableSymbols}
      value={selectedSymbols}
      onChange={(_, newValue) => onChange(newValue)}
      getOptionLabel={(option: SymbolInfo) => `${option.symbol} - ${option.name}`}
      isOptionEqualToValue={(option, value) => option.coinId === value.coinId}
      renderInput={(params) => (
        <TextField {...params} label="Currencies" placeholder="Search currencies..." />
      )}
      ChipProps={{ size: 'small' }}
    />
  )
}
