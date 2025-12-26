import { Input } from '@heroui/react'
import { useSettingsApiKey } from '../states/useSettingsApiKey'

export function SettingsApiKey() {
  const { value, onChange } = useSettingsApiKey()

  return (
    <Input
      label="API Key"
      type="password"
      placeholder="Enter API key"
      value={value}
      onValueChange={onChange}
    />
  )
}
