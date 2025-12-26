import { Select, SelectItem } from '@heroui/react'
import { REFRESH_OPTIONS } from '../constants/refreshOptions'
import { useSettingsRefreshInterval } from '../states/useSettingsRefreshInterval'

export function SettingsRefreshInterval() {
  const { refreshSeconds, onChange } = useSettingsRefreshInterval()

  return (
    <Select
      label="Refresh Interval"
      placeholder="Select interval"
      selectedKeys={[String(refreshSeconds)]}
      onSelectionChange={onChange}
    >
      {REFRESH_OPTIONS.map((opt) => (
        <SelectItem key={String(opt.value)}>{opt.label}</SelectItem>
      ))}
    </Select>
  )
}
