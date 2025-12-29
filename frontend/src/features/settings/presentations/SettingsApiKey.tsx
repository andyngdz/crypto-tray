import { Input, Label, TextField } from '@heroui/react'
import { useSettingsApiKey } from '@/features/settings/states/useSettingsApiKey'

export function SettingsApiKey() {
  const { value, onChange } = useSettingsApiKey()

  return (
    <TextField className="w-full">
      <Label>API Key</Label>
      <Input
        type="password"
        placeholder="Enter API key"
        value={value}
        onChange={(e) => onChange(e.target.value)}
      />
    </TextField>
  )
}
