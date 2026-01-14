import { FormControl, InputLabel, MenuItem, Select } from '@mui/material'

import { useSettingsProvider } from '@/features/settings/states/useSettingsProvider'

export function SettingsProvider() {
  const { providerId, providers, onChange } = useSettingsProvider()

  return (
    <FormControl fullWidth>
      <InputLabel>API Provider</InputLabel>
      <Select
        label="API Provider"
        value={providerId}
        onChange={(e) => onChange(e.target.value)}
      >
        {providers.map((p) => (
          <MenuItem key={p.id} value={p.id}>
            {p.name}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  )
}
