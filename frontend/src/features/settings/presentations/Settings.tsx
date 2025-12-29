import { SettingsView } from '@/features/settings/presentations/SettingsView'
import { StatusDisplay } from '@/features/settings/presentations/StatusDisplay'
import { useConfig } from '@/features/settings/states/useConfig'

export function Settings() {
  const { config, loading } = useConfig()

  if (loading) {
    return <StatusDisplay type="loading" />
  }

  if (!config) {
    return <StatusDisplay type="error" />
  }

  return <SettingsView />
}
