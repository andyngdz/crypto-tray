import { Checkbox, FormControlLabel } from '@mui/material'

import { useSettingsAutoStart } from '@/features/settings/states/useSettingsAutoStart'

export function SettingsAutoStart() {
  const { autoStart, onChange } = useSettingsAutoStart()

  return (
    <FormControlLabel
      control={
        <Checkbox
          checked={autoStart}
          onChange={(e) => onChange(e.target.checked)}
        />
      }
      label="Start on system startup"
    />
  )
}
