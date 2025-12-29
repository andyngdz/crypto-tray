import { useSettingsProvider } from '@/features/settings/states/useSettingsProvider'
import { Label, ListBox, Select } from '@heroui/react'

export function SettingsProvider() {
  const { providerId, providers, onChange } = useSettingsProvider()

  return (
    <Select
      className="w-full"
      placeholder="Select a provider"
      value={providerId}
      onChange={onChange}
    >
      <Label>API Provider</Label>
      <Select.Trigger>
        <Select.Value />
        <Select.Indicator />
      </Select.Trigger>
      <Select.Popover>
        <ListBox>
          {providers.map((p) => (
            <ListBox.Item key={p.id} id={p.id} textValue={p.name}>
              {p.name}
              <ListBox.ItemIndicator />
            </ListBox.Item>
          ))}
        </ListBox>
      </Select.Popover>
    </Select>
  )
}
