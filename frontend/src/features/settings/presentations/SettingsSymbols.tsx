import { Label, ListBox, Select } from '@heroui/react'

import { useConfig } from '@/features/settings/states/useConfig'
import { useSettingsSymbols } from '@/features/settings/states/useSettingsSymbols'

export function SettingsSymbols() {
  const { availableSymbols } = useConfig()
  const { value, onChange } = useSettingsSymbols()

  return (
    <Select
      className="w-full"
      placeholder="Select currencies"
      selectionMode="multiple"
      value={value}
      onChange={onChange}
    >
      <Label>Currencies</Label>
      <Select.Trigger>
        <Select.Value />
        <Select.Indicator />
      </Select.Trigger>
      <Select.Popover>
        <ListBox selectionMode="multiple">
          {availableSymbols.map((s) => (
            <ListBox.Item key={s.id} id={s.id} textValue={s.name}>
              {s.id.toUpperCase()}, {s.name}
              <ListBox.ItemIndicator />
            </ListBox.Item>
          ))}
        </ListBox>
      </Select.Popover>
    </Select>
  )
}
