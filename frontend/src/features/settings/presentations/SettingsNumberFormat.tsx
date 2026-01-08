import { FormControl, InputLabel, MenuItem, Select } from '@mui/material'

import { NUMBER_FORMAT_OPTIONS } from '@/features/settings/constants/numberFormatOptions'
import { useSettingsNumberFormat } from '@/features/settings/states/useSettingsNumberFormat'

export function SettingsNumberFormat() {
  const { numberFormat, onChange } = useSettingsNumberFormat()

  return (
    <FormControl fullWidth>
      <InputLabel>Number Format</InputLabel>
      <Select
        label="Number Format"
        value={numberFormat}
        onChange={(e) => onChange(e.target.value)}
      >
        {NUMBER_FORMAT_OPTIONS.map((opt) => (
          <MenuItem key={opt.value} value={opt.value}>
            {opt.label}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  )
}
