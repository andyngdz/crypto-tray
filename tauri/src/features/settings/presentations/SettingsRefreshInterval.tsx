import { FormControl, InputLabel, MenuItem, Select } from '@mui/material'

import { REFRESH_OPTIONS } from '@/features/settings/constants/refreshOptions'
import { useSettingsRefreshInterval } from '@/features/settings/states/useSettingsRefreshInterval'

export function SettingsRefreshInterval() {
  const { refreshSeconds, onChange } = useSettingsRefreshInterval()

  return (
    <FormControl fullWidth>
      <InputLabel>Refresh Interval</InputLabel>
      <Select
        label="Refresh Interval"
        value={String(refreshSeconds)}
        onChange={(e) => onChange(e.target.value)}
      >
        {REFRESH_OPTIONS.map((opt) => (
          <MenuItem key={opt.value} value={String(opt.value)}>
            {opt.label}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  )
}
