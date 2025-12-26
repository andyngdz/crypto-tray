import { Select, SelectItem } from '@heroui/react'
import { useSettingsProvider } from '../states/useSettingsProvider'

export function SettingsProvider() {
  const { providerId, providers, onChange } = useSettingsProvider()

  return (
    <Select
      label="API Provider"
      placeholder="Select a provider"
      selectedKeys={[providerId]}
      onSelectionChange={onChange}
    >
      {providers.map((p) => (
        <SelectItem key={p.id}>{p.name}</SelectItem>
      ))}
    </Select>
  )
}
