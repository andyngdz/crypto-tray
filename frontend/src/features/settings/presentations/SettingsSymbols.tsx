import {
  Autocomplete,
  Chip,
  ListItem,
  ListItemIcon,
  ListItemText,
  TextField,
} from '@mui/material'

import { CryptoIcon } from '@/components/CryptoIcon'
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
      getOptionLabel={(option: SymbolInfo) =>
        `${option.symbol} - ${option.name}`
      }
      isOptionEqualToValue={(option, value) => option.coinId === value.coinId}
      renderOption={(props, option) => {
        const { key, ...otherProps } = props
        return (
          <ListItem key={key} disablePadding {...otherProps}>
            <ListItemIcon sx={{ minWidth: 36 }}>
              <CryptoIcon symbol={option.symbol} />
            </ListItemIcon>
            <ListItemText primary={`${option.symbol} - ${option.name}`} />
          </ListItem>
        )
      }}
      renderValue={(value, getItemProps) =>
        value.map((option, index) => {
          const { key, ...itemProps } = getItemProps({ index })
          return (
            <Chip
              key={key}
              avatar={<CryptoIcon symbol={option.symbol} size={16} />}
              label={option.symbol}
              {...itemProps}
            />
          )
        })
      }
      renderInput={(params) => (
        <TextField
          {...params}
          label="Currencies"
          placeholder="Search currencies..."
        />
      )}
    />
  )
}
