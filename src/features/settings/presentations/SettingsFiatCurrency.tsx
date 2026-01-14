import { FormControl, InputLabel, MenuItem, Select } from '@mui/material'

import { FIAT_CURRENCY_OPTIONS } from '@/features/settings/constants/fiatCurrencyOptions'
import { useSettingsFiatCurrency } from '@/features/settings/states/useSettingsFiatCurrency'

export function SettingsFiatCurrency() {
  const { displayCurrency, onChange } = useSettingsFiatCurrency()

  return (
    <FormControl fullWidth>
      <InputLabel>Display Currency</InputLabel>
      <Select
        label="Display Currency"
        value={displayCurrency}
        onChange={(e) => onChange(e.target.value)}
      >
        {FIAT_CURRENCY_OPTIONS.map((opt) => (
          <MenuItem key={opt.value} value={opt.value}>
            {opt.label}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  )
}
