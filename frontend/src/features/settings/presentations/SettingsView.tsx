import { SettingsApiKey } from '@/features/settings/presentations/SettingsApiKey'
import { SettingsProvider } from '@/features/settings/presentations/SettingsProvider'
import { SettingsRefreshInterval } from '@/features/settings/presentations/SettingsRefreshInterval'
import { SettingsSymbols } from '@/features/settings/presentations/SettingsSymbols'
import { hideWindow } from '@/features/settings/services/configService'
import { useConfig } from '@/features/settings/states/useConfig'
import { useSettingsProvider } from '@/features/settings/states/useSettingsProvider'
import { Button } from '@heroui/react'

export function SettingsView() {
  const { saving, error } = useConfig()
  const { currentProvider } = useSettingsProvider()

  return (
    <div className="p-6 min-h-screen flex flex-col gap-4">
      {error && (
        <div className="p-3 text-sm text-white bg-danger rounded-lg">
          {error}
        </div>
      )}

      <SettingsProvider />

      {currentProvider?.requiresApiKey && <SettingsApiKey />}

      <SettingsSymbols />

      <SettingsRefreshInterval />

      <Button
        className="w-full mt-auto"
        onPress={hideWindow}
        isDisabled={saving}
      >
        {saving ? 'Saving...' : 'Close'}
      </Button>
    </div>
  )
}
