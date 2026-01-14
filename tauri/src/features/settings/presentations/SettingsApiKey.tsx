import { TextField } from '@mui/material'

import { useSettingsApiKey } from '@/features/settings/states/useSettingsApiKey'

export function SettingsApiKey() {
  const { value, onChange } = useSettingsApiKey()

  return (
    <TextField
      fullWidth
      label="API Key"
      type="password"
      placeholder="Enter API key"
      value={value}
      onChange={(e) => onChange(e.target.value)}
    />
  )
}
