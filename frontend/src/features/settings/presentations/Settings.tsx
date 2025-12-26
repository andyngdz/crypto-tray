import { useConfig } from '../states/useConfig'
import { SettingsView } from './SettingsView'
import { StatusDisplay } from './StatusDisplay'

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
