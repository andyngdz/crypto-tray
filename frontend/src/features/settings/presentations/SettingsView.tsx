import { Button } from '@heroui/react'
import { hideWindow } from '../services/configService'
import { useConfig } from '../states/useConfig'
import { useSettingsProvider } from '../states/useSettingsProvider'
import { SettingsApiKey } from './SettingsApiKey'
import { SettingsProvider } from './SettingsProvider'
import { SettingsRefreshInterval } from './SettingsRefreshInterval'

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
