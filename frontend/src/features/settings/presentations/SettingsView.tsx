import { Alert, Box, Button } from '@mui/material'

import { SettingsApiKey } from '@/features/settings/presentations/SettingsApiKey'
import { SettingsFiatCurrency } from '@/features/settings/presentations/SettingsFiatCurrency'
import { SettingsNumberFormat } from '@/features/settings/presentations/SettingsNumberFormat'
import { SettingsProvider } from '@/features/settings/presentations/SettingsProvider'
import { SettingsRefreshInterval } from '@/features/settings/presentations/SettingsRefreshInterval'
import { SettingsSection } from '@/features/settings/presentations/SettingsSection'
import { SettingsSymbols } from '@/features/settings/presentations/SettingsSymbols'
import { hideWindow } from '@/features/settings/services/configService'
import { useConfig } from '@/features/settings/states/useConfig'
import { useSettingsProvider } from '@/features/settings/states/useSettingsProvider'

export function SettingsView() {
  const { saving, error } = useConfig()
  const { currentProvider } = useSettingsProvider()

  return (
    <Box sx={{ p: 3, minHeight: 'calc(100vh - 32px)', display: 'flex', flexDirection: 'column', gap: 2 }}>
      {error && <Alert severity="error">{error}</Alert>}

      <SettingsSection title="General">
        <SettingsProvider />

        {currentProvider?.requiresApiKey && <SettingsApiKey />}

        <SettingsSymbols />

        <SettingsFiatCurrency />

        <SettingsRefreshInterval />
      </SettingsSection>

      <SettingsSection title="Formatting">
        <SettingsNumberFormat />
      </SettingsSection>

      <Button
        variant="contained"
        fullWidth
        onClick={hideWindow}
        disabled={saving}
        sx={{ mt: 'auto' }}
      >
        {saving ? 'Saving...' : 'Close'}
      </Button>
    </Box>
  )
}
