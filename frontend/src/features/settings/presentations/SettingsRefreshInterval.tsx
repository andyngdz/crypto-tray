import { REFRESH_OPTIONS } from '@/features/settings/constants/refreshOptions'
import { useSettingsRefreshInterval } from '@/features/settings/states/useSettingsRefreshInterval'
import { Label, ListBox, Select } from '@heroui/react'

export function SettingsRefreshInterval() {
  const { refreshSeconds, onChange } = useSettingsRefreshInterval()

  return (
    <Select
      className="w-full"
      placeholder="Select interval"
      value={String(refreshSeconds)}
      onChange={onChange}
    >
      <Label>Refresh Interval</Label>
      <Select.Trigger>
        <Select.Value />
        <Select.Indicator />
      </Select.Trigger>
      <Select.Popover>
        <ListBox>
          {REFRESH_OPTIONS.map((opt) => (
            <ListBox.Item
              key={String(opt.value)}
              id={String(opt.value)}
              textValue={opt.label}
            >
              {opt.label}
              <ListBox.ItemIndicator />
            </ListBox.Item>
          ))}
        </ListBox>
      </Select.Popover>
    </Select>
  )
}
