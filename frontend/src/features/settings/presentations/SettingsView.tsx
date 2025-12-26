import {
  Button,
  Card,
  CardBody,
  CardHeader,
  Divider,
} from '@heroui/react'
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
    <div className="min-h-screen p-6 bg-background">
      <Card className="max-w-md mx-auto">
        <CardHeader className="flex flex-col items-start gap-1">
          <h1 className="text-xl font-bold">Crypto Tray Settings</h1>
          <p className="text-small text-default-500">
            Configure your price tracker
          </p>
        </CardHeader>
        <Divider />
        <CardBody className="gap-4">
          {error && (
            <div className="p-3 text-sm text-white bg-danger rounded-lg">
              {error}
            </div>
          )}

          <SettingsProvider />

          {currentProvider?.requiresApiKey && <SettingsApiKey />}

          <SettingsRefreshInterval />

          <Divider className="my-2" />

          <Button
            color="primary"
            className="w-full"
            onPress={hideWindow}
            isLoading={saving}
          >
            {saving ? 'Saving...' : 'Close'}
          </Button>
        </CardBody>
      </Card>
    </div>
  )
}
