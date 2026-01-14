import { Alert, Box, Button, Divider } from '@mui/material'

import { SettingsApiKey } from '@/features/settings/presentations/SettingsApiKey'
import { SettingsAutoStart } from '@/features/settings/presentations/SettingsAutoStart'
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
    <Box
      sx={{
        minHeight: 'calc(100vh - 32px)',
        display: 'flex',
        flexDirection: 'column',
      }}
    >
      {error && <Alert severity="error">{error}</Alert>}

      <SettingsSection title="General">
        <SettingsProvider />

        {currentProvider?.requiresApiKey && <SettingsApiKey />}

        <SettingsSymbols />

        <SettingsFiatCurrency />

        <SettingsRefreshInterval />
      </SettingsSection>
      <Divider sx={{ opacity: 0.5 }} />
      <SettingsSection title="Formatting">
        <SettingsNumberFormat />
      </SettingsSection>
      <Divider sx={{ opacity: 0.5 }} />
      <SettingsSection title="System">
        <SettingsAutoStart />
      </SettingsSection>
      <Divider sx={{ opacity: 0.5 }} />
      <Button
        variant="contained"
        onClick={hideWindow}
        disabled={saving}
        sx={{ mt: 'auto' }}
        fullWidth
      >
        {saving ? 'Saving...' : 'Close'}
      </Button>
    </Box>
  )
}
